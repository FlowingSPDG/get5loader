package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/hydrogen18/stalecucumber"
	"math/rand"
	"net"
	"os/exec"
	"strings"

	"github.com/Acidic9/steam"
	gosteam "github.com/kidoman/go-steam"
	"strconv"
)

// Config Configration Struct for config.ini
type Config struct {
	SteamAPIKey string
	DefaultPage string
	SQLHost     string
	SQLUser     string
	SQLPass     string
	SQLPort     int
	SQLDBName   string
	HOST        string
}

var (
	Cnf Config
)

func init() {
	c, _ := ini.Load("config.ini")
	Cnf = Config{
		SteamAPIKey: c.Section("Steam").Key("APIKey").MustString(""),
		DefaultPage: c.Section("GET5").Key("DefaultPage").MustString(""),
		HOST:        c.Section("GET5").Key("HOST").MustString(""),
		SQLHost:     c.Section("sql").Key("host").MustString(""),
		SQLUser:     c.Section("sql").Key("user").MustString(""),
		SQLPass:     c.Section("sql").Key("pass").MustString(""),
		SQLPort:     c.Section("sql").Key("port").MustInt(3306),
		SQLDBName:   c.Section("sql").Key("database").MustString(""),
	}
}

func checkIP(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", ip)
		return false
	}
	return true
}

// FormatMapName Formats correct map name.
func FormatMapName(mapname string) string {
	FormattedNames := make(map[string]string)
	FormattedNames["de_dust2"] = "Dust II"
	FormattedNames["de_mirage"] = "Mirage"
	FormattedNames["de_overpass"] = "Overpass"
	FormattedNames["de_vertigo"] = "Vertigo"
	FormattedNames["de_nuke"] = "NUKE"
	FormattedNames["de_train"] = "Train"
	FormattedNames["de_inferno"] = "Inferno"
	FormattedNames["de_cache"] = "Cache"
	FormattedNames["de_cbble"] = "Cobblestone"

	if _, ok := FormattedNames[mapname]; ok {
		return FormattedNames[mapname]
	}
	return ""
}

// SendRCON Sends Remote-Commands to specific IP SRCDS.
func SendRCON(host string, port int, pass string, cmd string) (string, error) {
	if !checkIP(host) {
		return "", fmt.Errorf("Specified IP is not valid")
	}
	dest := host + ":" + strconv.Itoa(port)
	o := &gosteam.ConnectOptions{RCONPassword: pass}
	rcon, err := gosteam.Connect(dest, o)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer rcon.Close()

	resp, err := rcon.Send(cmd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp, nil
}

// CheckServerConnection Check server pulse by sending "status" command
func CheckServerConnection(IPString string, Port int, RconPassword string) bool {
	_, err := SendRCON(IPString, Port, RconPassword, "status")
	if err != nil {
		return false
	}
	return true
}

// GET5AvailableDatas Struct for Get5 availability check.
type GET5AvailableDatas struct {
	Gamestate     int    `json:"gamestate"`
	Available     int    `json:"available"`
	PluginVersion string `json:"plugin_version"`
}

// CheckServerAvailability if server is usable for get5_web
func CheckServerAvailability(IPString string, Port int, RconPassword string) (GET5AvailableDatas, error) { // available or error string
	var data = GET5AvailableDatas{}
	resp, err := SendRCON(IPString, Port, RconPassword, "get5_web_avaliable")
	if err != nil {
		return data, fmt.Errorf("Connect fails")
	}
	jsonBytes := ([]byte)(resp)
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return data, fmt.Errorf("Error reading get5_web_avaliable response")
	}
	if strings.Contains(resp, "Unknown command") {
		return data, fmt.Errorf("Either get5 or get5_apistats plugin missin")
	}
	if data.Gamestate != 0 {
		return data, fmt.Errorf("Server already has a get5 match setup")
	}
	return data, nil

}

// GetVersion Gets get5-web-go version from github
func GetVersion() (string, error) {
	//root_dir, err := os.Executable()
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "", nil
	}
	return string(out), nil
}

// SteamID64sToPickle Pickles Steamid64 array
func SteamID64sToPickle(ids []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := stalecucumber.NewPickler(buf).Pickle(ids)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PickleToSteamID64s Un-pickles Python array to SteamID64 array
func PickleToSteamID64s(Pickles []byte) ([]string, error) {
	reader := bytes.NewReader(Pickles)
	res := make([]string, 0)
	err := stalecucumber.UnpackInto(&res).From(stalecucumber.Unpickle(reader))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandString Generates random string
func RandString(n int) string {
	const rs2Letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = rs2Letters[rand.Intn(len(rs2Letters))]
	}
	return string(b)
}

// AuthToSteamID64 Converts Auth to SteamID64
func AuthToSteamID64(auth string) (string, error) {
	auth = strings.TrimSpace(auth)
	if strings.Contains(auth, "steamcommunity.com/id/") {
		s := steam.SearchForID(auth, Cnf.SteamAPIKey)
		if s == 0 {
			return "", fmt.Errorf("User not found")
		}
		return strconv.Itoa(int(s)), nil
	} else if strings.Contains(auth, "steamcommunity.com/profiles/") {
		s := steam.SearchForID(auth, Cnf.SteamAPIKey)
		if s == 0 {
			return "", fmt.Errorf("User not found")
		}
		return string(s), nil
	} else if strings.HasPrefix(auth, "1:0") || strings.HasPrefix(auth, "1:1") {
		var s steam.SteamID = steam.SteamID("STEAM_" + auth)
		return strconv.Itoa(int(steam.SteamIDToSteamID64(s))), nil
	} else if strings.HasPrefix(auth, "STEAM_") {
		var s steam.SteamID = steam.SteamID(auth)
		return strconv.Itoa(int(steam.SteamIDToSteamID64(s))), nil
	} else if strings.HasPrefix(auth, "7656119") && !strings.Contains(auth, "steam") {
		return auth, nil
	} else if strings.HasPrefix(auth, "[U:1:") {
		var s steam.SteamID3 = steam.SteamID3(auth)
		return strconv.Itoa(int(steam.SteamID3ToSteamID64(s))), nil
	}
	s := steam.SearchForID(auth, Cnf.SteamAPIKey)
	if s == 0 {
		return "", fmt.Errorf("User not found")
	}
	return strconv.Itoa(int(s)), nil
}

// IsValidSteamID Checks if steamid is correct format or not
func IsValidSteamID(auth string) bool {
	_, err := AuthToSteamID64(auth)
	if err != nil {
		return false
	}
	return true
}
