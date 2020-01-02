/**
 * =============================================================================
 * Get5 web API integration
 * Copyright (C) 2016. Sean Lewis.  All rights reserved.
 * =============================================================================
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

#include "include/get5.inc"
#include "include/logdebug.inc"
#include <cstrike>
#include <sourcemod>
#include "get5/util.sp"
#include "get5/version.sp"

#include <SteamWorks>
#include <system2> // github.com/dordnung/System2/
#include <json> // github.com/clugg/sm-json
#include "get5/jsonhelpers.sp"

#pragma semicolon 1
#pragma newdecls required

int g_MatchID = -1;

ConVar g_APIKeyCvar;
char g_APIKey[128];

ConVar g_APIURLCvar;
char g_APIURL[128];

char g_storedAPIURL[128];
char g_storedAPIKey[128];

ConVar g_FTPHostCvar;
char g_FTPHost[128];

ConVar g_FTPUsernameCvar;
char g_FTPUsername[128];

ConVar g_FTPPasswordCvar;
char g_FTPPassword[128];

ConVar g_FTPPortCvar;
int g_FTPPort;

ConVar g_FTPEnableCvar;
bool g_FTPEnable;


#define LOGO_DIR "resource/flash/econ/tournaments/teams"
#define PANO_DIR "materials/panorama/images/tournaments/teams"
// clang-format off
public Plugin myinfo = {
  name = "Get5 Web API Integration",
  author = "splewis/phlexplexico",
  description = "Records match stats to a get5-web api",
  version = "1.0",
  url = "https://github.com/phlexplexico/get5-web"
};
// clang-format on

public void OnPluginStart() {
  InitDebugLog("get5_debug", "get5_api");
  LogDebug("OnPluginStart version=%s", PLUGIN_VERSION);

  g_FTPHostCvar = 
      CreateConVar("get5_api_ftp_host", "ftp://example.com", "Remote FTP Host. Make sure you do NOT have the trailing slash. Include the path to the directory you wish to have.", FCVAR_PROTECTED);

  g_FTPPortCvar = 
      CreateConVar("get5_api_ftp_port", "21", "Remote FTP Port", FCVAR_PROTECTED);

  g_FTPUsernameCvar =
      CreateConVar("get5_api_ftp_username", "username", "Username for the FTP connection.", FCVAR_PROTECTED);

  g_FTPPasswordCvar = 
      CreateConVar("get5_api_ftp_password", "supersecret", "Password for the FTP user. Leave blank if no password.", FCVAR_PROTECTED);

  g_FTPEnableCvar = 
      CreateConVar("get5_api_ftp_enabled", "0", "0 Disables FTP Upload, 1 Enables.");

  g_APIKeyCvar =
      CreateConVar("get5_web_api_key", "", "Match API key, this is automatically set through rcon", FCVAR_DONTRECORD);
  HookConVarChange(g_APIKeyCvar, ApiInfoChanged);

  g_APIURLCvar = CreateConVar("get5_web_api_url", "", "URL the get5 api is hosted at, IGNORE AS IT IS SYSTEM SET.", FCVAR_DONTRECORD);

  HookConVarChange(g_APIURLCvar, ApiInfoChanged);

  RegConsoleCmd("get5_web_avaliable",
                Command_Avaliable);  // legacy version since I'm bad at spelling
  RegConsoleCmd("get5_web_available", Command_Avaliable);
  /** Create and exec plugin's configuration file **/
  AutoExecConfig(true, "get5api");
  
}

public Action Command_Avaliable(int client, int args) {
  char versionString[64] = "unknown";
  ConVar versionCvar = FindConVar("get5_version");
  if (versionCvar != null) {
    versionCvar.GetString(versionString, sizeof(versionString));
  }

  JSON_Object json = new JSON_Object();

  json.SetInt("gamestate", view_as<int>(Get5_GetGameState()));
  json.SetInt("avaliable", 1); // legacy version since I'm bad at spelling
  json.SetInt("available", 1);
  json.SetString("plugin_version", versionString);

  char buffer[128];
  json.Encode(buffer, sizeof(buffer));
  ReplyToCommand(client, buffer);

  delete json;

  return Plugin_Handled;
}

public void ApiInfoChanged(ConVar convar, const char[] oldValue, const char[] newValue) {
  g_APIKeyCvar.GetString(g_APIKey, sizeof(g_APIKey));
  g_APIURLCvar.GetString(g_APIURL, sizeof(g_APIURL));

  // Add a trailing backslash to the api url if one is missing.
  int len = strlen(g_APIURL);
  if (len > 0 && g_APIURL[len - 1] != '/') {
    StrCat(g_APIURL, sizeof(g_APIURL), "/");
  }

  LogDebug("get5_web_api_url now set to %s", g_APIURL);
}

static Handle CreateRequest(EHTTPMethod httpMethod, const char[] apiMethod, any:...) {
  char url[1024];
  Format(url, sizeof(url), "%s%s", g_APIURL, apiMethod);

  char formattedUrl[1024];
  VFormat(formattedUrl, sizeof(formattedUrl), url, 3);

  LogDebug("Trying to create request to url %s", formattedUrl);

  Handle req = SteamWorks_CreateHTTPRequest(httpMethod, formattedUrl);
  if (StrEqual(g_APIKey, "")) {
    // Not using a web interface.
    return INVALID_HANDLE;

  } else if (req == INVALID_HANDLE) {
    LogError("Failed to create request to %s", formattedUrl);
    return INVALID_HANDLE;

  } else {
    SteamWorks_SetHTTPCallbacks(req, RequestCallback);
    AddStringParam(req, "key", g_APIKey);
    return req;
  }
}

static Handle CreateDemoRequest(EHTTPMethod httpMethod, const char[] apiMethod, any:...) {
  char url[1024];
  Format(url, sizeof(url), "%s%s", g_storedAPIURL, apiMethod);

  char formattedUrl[1024];
  VFormat(formattedUrl, sizeof(formattedUrl), url, 3);

  LogDebug("Trying to create request to url %s", formattedUrl);

  Handle req = SteamWorks_CreateHTTPRequest(httpMethod, formattedUrl);
  if (StrEqual(g_storedAPIKey, "")) {
    // Not using a web interface.
    return INVALID_HANDLE;

  } else if (req == INVALID_HANDLE) {
    LogError("Failed to create request to %s", formattedUrl);
    return INVALID_HANDLE;

  } else {
    SteamWorks_SetHTTPCallbacks(req, RequestCallback);
    AddStringParam(req, "key", g_APIKey);
    return req;
  }
}

public int RequestCallback(Handle request, bool failure, bool requestSuccessful,
                    EHTTPStatusCode statusCode) {
  if (failure || !requestSuccessful) {
    LogError("API request failed, HTTP status code = %d", statusCode);
    char response[1024];
    SteamWorks_GetHTTPResponseBodyData(request, response, sizeof(response));
    LogError(response);
    return;
  }
}

public void Get5_OnBackupRestore() {
  char matchid[64];
  Get5_GetMatchID(matchid, sizeof(matchid));
  g_MatchID = StringToInt(matchid);
}

public void Get5_OnSeriesInit() {
  char matchid[64];
  Get5_GetMatchID(matchid, sizeof(matchid));
  g_MatchID = StringToInt(matchid);

  // Handle new logos.
  if (!DirExists(LOGO_DIR)) {
    if (!CreateDirectory(LOGO_DIR, 755)) {
      LogError("Failed to create logo directory: %s", LOGO_DIR);
    }
  }
  if (!DirExists(PANO_DIR)) {
    if (!CreateDirectory(PANO_DIR, 755)) {
      LogError("Failed to create logo directory: %s", PANO_DIR);
    }
  }

  char logo1[32];
  char logo2[32];
  GetConVarStringSafe("mp_teamlogo_1", logo1, sizeof(logo1));
  GetConVarStringSafe("mp_teamlogo_2", logo2, sizeof(logo2));
  CheckForLogo(logo1);
  CheckForLogo(logo2);
}

public void CheckForLogo(const char[] logo) {
  if (StrEqual(logo, "")) {
    return;
  }

  char logoPath[PLATFORM_MAX_PATH + 1];
  char svgLogoPath[PLATFORM_MAX_PATH +1];
  Format(logoPath, sizeof(logoPath), "%s/%s.png", LOGO_DIR, logo);
  Format(svgLogoPath, sizeof(svgLogoPath), "%s/%s.svg", PANO_DIR, logo);

  // Try to fetch the file if we don't have it.
  if (!FileExists(logoPath)) {
    LogDebug("Fetching logo for %s", logo);
    Handle req = CreateRequest(k_EHTTPMethodGET, "/static/resource/csgo/resource/flash/econ/tournaments/teams/%s.png", logo);
    if (req == INVALID_HANDLE) {
      return;
    }

    Handle pack = CreateDataPack();
    WritePackString(pack, logo);

    SteamWorks_SetHTTPRequestContextValue(req, view_as<int>(pack));
    SteamWorks_SetHTTPCallbacks(req, LogoCallback);
    SteamWorks_SendHTTPRequest(req);
  }

  //Attempt to get SVG.
  if (!FileExists(svgLogoPath)) {
    LogDebug("Fetching logo for %s", logo);
    Handle req = CreateRequest(k_EHTTPMethodGET, "/static/resource/csgo/materials/panorama/images/tournaments/teams/%s.svg", logo);
    if (req == INVALID_HANDLE) {
      return;
    }

    Handle svgPack = CreateDataPack();
    WritePackString(svgPack, logo);

    SteamWorks_SetHTTPRequestContextValue(req, view_as<int>(svgPack));
    SteamWorks_SetHTTPCallbacks(req, LogoCallbackSvg);
    SteamWorks_SendHTTPRequest(req);
  }
}

public int LogoCallback(Handle request, bool failure, bool successful, EHTTPStatusCode status, int data) {
  if (failure || !successful) {
    LogError("Logo request failed, status code = %d", status);
    return;
  }

  DataPack pack = view_as<DataPack>(data);
  pack.Reset();
  char logo[32];
  pack.ReadString(logo, sizeof(logo));

  char logoPath[PLATFORM_MAX_PATH + 1];
  Format(logoPath, sizeof(logoPath), "%s/%s.png", LOGO_DIR, logo);

  LogMessage("Saved logo for %s to %s", logo, logoPath);
  SteamWorks_WriteHTTPResponseBodyToFile(request, logoPath);
}

public int LogoCallbackSvg(Handle request, bool failure, bool successful, EHTTPStatusCode status, int data) {
  if (failure || !successful) {
    LogError("Logo request failed, status code = %d", status);
    return;
  }

  DataPack pack = view_as<DataPack>(data);
  pack.Reset();
  char logo[32];
  pack.ReadString(logo, sizeof(logo));

  char svgLogoPath[PLATFORM_MAX_PATH + 1];
  Format(svgLogoPath, sizeof(svgLogoPath), "%s/%s.svg", PANO_DIR, logo);

  LogMessage("Saved logo for %s to %s", logo, svgLogoPath);
  SteamWorks_WriteHTTPResponseBodyToFile(request, svgLogoPath);
}

public void Get5_OnGoingLive(int mapNumber) {
  char mapName[64];
  g_FTPEnable = g_FTPEnableCvar.BoolValue;
  
  GetCurrentMap(mapName, sizeof(mapName));
  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/map/%d/start", g_MatchID, mapNumber);
  if (req != INVALID_HANDLE) {
    AddStringParam(req, "mapname", mapName);
    SteamWorks_SendHTTPRequest(req);
  }
  // Store Cvar since it gets reset after match finishes?
  if (g_FTPEnable) {
    Format(g_storedAPIKey, sizeof(g_storedAPIKey), g_APIKey);
    Format(g_storedAPIURL, sizeof(g_storedAPIURL), g_APIURL);
  }
  Get5_AddLiveCvar("get5_web_api_key", g_APIKey);
  Get5_AddLiveCvar("get5_web_api_url", g_APIURL);
  
}

public void UpdateRoundStats(int mapNumber) {
  int t1score = CS_GetTeamScore(Get5_MatchTeamToCSTeam(MatchTeam_Team1));
  int t2score = CS_GetTeamScore(Get5_MatchTeamToCSTeam(MatchTeam_Team2));

  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/map/%d/update", g_MatchID, mapNumber);
  if (req != INVALID_HANDLE) {
    AddIntParam(req, "team1score", t1score);
    AddIntParam(req, "team2score", t2score);
    SteamWorks_SendHTTPRequest(req);
  }

  // Update player stats
  KeyValues kv = new KeyValues("Stats");
  Get5_GetMatchStats(kv);
  char mapKey[32];
  Format(mapKey, sizeof(mapKey), "map%d", mapNumber);
  if (kv.JumpToKey(mapKey)) {
    if (kv.JumpToKey("team1")) {
      UpdatePlayerStats(kv, MatchTeam_Team1);
      kv.GoBack();
    }
    if (kv.JumpToKey("team2")) {
      UpdatePlayerStats(kv, MatchTeam_Team2);
      kv.GoBack();
    }
    kv.GoBack();
  }
  delete kv;
}

public void Get5_OnMapResult(const char[] map, MatchTeam mapWinner, int team1Score, int team2Score,
                      int mapNumber) {
  char winnerString[64];
  GetTeamString(mapWinner, winnerString, sizeof(winnerString));

  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/map/%d/finish", g_MatchID, mapNumber);
  if (req != INVALID_HANDLE) {
    AddIntParam(req, "team1score", team1Score);
    AddIntParam(req, "team2score", team2Score);
    AddStringParam(req, "winner", winnerString);
    SteamWorks_SendHTTPRequest(req);
  }
}



static void AddIntStat(Handle req, KeyValues kv, const char[] field) {
  AddIntParam(req, field, kv.GetNum(field));
}

public void UpdatePlayerStats(KeyValues kv, MatchTeam team) {
  char name[MAX_NAME_LENGTH];
  char auth[AUTH_LENGTH];
  int mapNumber = MapNumber();

  if (kv.GotoFirstSubKey()) {
    do {
      kv.GetSectionName(auth, sizeof(auth));
      kv.GetString("name", name, sizeof(name));
      char teamString[16];
      GetTeamString(team, teamString, sizeof(teamString));

      Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/map/%d/player/%s/update", g_MatchID,
                                 mapNumber, auth);
      if (req != INVALID_HANDLE) {
        AddStringParam(req, "team", teamString);
        AddStringParam(req, "name", name);
        AddIntStat(req, kv, STAT_KILLS);
        AddIntStat(req, kv, STAT_DEATHS);
        AddIntStat(req, kv, STAT_ASSISTS);
        AddIntStat(req, kv, STAT_FLASHBANG_ASSISTS);
        AddIntStat(req, kv, STAT_TEAMKILLS);
        AddIntStat(req, kv, STAT_SUICIDES);
        AddIntStat(req, kv, STAT_DAMAGE);
        AddIntStat(req, kv, STAT_HEADSHOT_KILLS);
        AddIntStat(req, kv, STAT_ROUNDSPLAYED);
        AddIntStat(req, kv, STAT_BOMBPLANTS);
        AddIntStat(req, kv, STAT_BOMBDEFUSES);
        AddIntStat(req, kv, STAT_1K);
        AddIntStat(req, kv, STAT_2K);
        AddIntStat(req, kv, STAT_3K);
        AddIntStat(req, kv, STAT_4K);
        AddIntStat(req, kv, STAT_5K);
        AddIntStat(req, kv, STAT_V1);
        AddIntStat(req, kv, STAT_V2);
        AddIntStat(req, kv, STAT_V3);
        AddIntStat(req, kv, STAT_V4);
        AddIntStat(req, kv, STAT_V5);
        AddIntStat(req, kv, STAT_FIRSTKILL_T);
        AddIntStat(req, kv, STAT_FIRSTKILL_CT);
        AddIntStat(req, kv, STAT_FIRSTDEATH_T);
        AddIntStat(req, kv, STAT_FIRSTDEATH_CT);
        AddIntStat(req, kv, STAT_TRADEKILL);
        SteamWorks_SendHTTPRequest(req);
      }

    } while (kv.GotoNextKey());
    kv.GoBack();
  }
}

static void AddStringParam(Handle request, const char[] key, const char[] value) {
  if (!SteamWorks_SetHTTPRequestGetOrPostParameter(request, key, value)) {
    LogError("Failed to add http param %s=%s", key, value);
  } else {
    LogDebug("Added param %s=%s to request", key, value);
  }
}

static void AddIntParam(Handle request, const char[] key, int value) {
  char buffer[32];
  IntToString(value, buffer, sizeof(buffer));
  AddStringParam(request, key, buffer);
}

public void Get5_OnMapVetoed(MatchTeam team, const char[] map){
  char teamString[64];
  GetTeamString(team, teamString, sizeof(teamString));
  LogDebug("Map Veto START team %s map vetoed %s", team, map);
  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/vetoUpdate", g_MatchID);
  if (req != INVALID_HANDLE) {
      AddStringParam(req, "map", map);
      AddStringParam(req, "teamString", teamString);
      AddStringParam(req, "pick_or_veto", "ban");
      SteamWorks_SendHTTPRequest(req);
  }
  LogDebug("Accepted Map Veto.");
}

public void Get5_OnDemoFinished(const char[] filename){
  g_FTPEnable = g_FTPEnableCvar.BoolValue;
  if (g_FTPEnable) {
    LogDebug("About to enter UploadDemo.");
    int mapNumber = MapNumber();
    char zippedFile[PLATFORM_MAX_PATH];
    char formattedURL[PLATFORM_MAX_PATH];
    UploadDemo(filename, zippedFile);

    Handle req = CreateDemoRequest(k_EHTTPMethodPOST, "match/%d/map/%d/demo", g_MatchID, mapNumber-1);
    LogDebug("Our api url: %s", g_storedAPIURL);
    // Send URL to store in database to show users at end of match.
    // This requires anonmyous downloads on the FTP server unless
    // you give out usernames.
    if (req != INVALID_HANDLE) {
        Format(formattedURL, sizeof(formattedURL), "%sstatic/demos/%s", g_storedAPIURL, zippedFile);
        LogDebug("Our URL: %s", formattedURL);
        AddStringParam(req, "demoFile", formattedURL);
        SteamWorks_SendHTTPRequest(req);
    }
    // Need to store as get5 recycles the configs before the demos finish recording.
    Format(g_storedAPIKey, sizeof(g_storedAPIKey), "");
    Format(g_storedAPIURL, sizeof(g_storedAPIURL), "");
  }
}

public void UploadDemo(const char[] filename, char zippedFile[PLATFORM_MAX_PATH]){
  char remoteDemoPath[PLATFORM_MAX_PATH];
  if(filename[0]){
    g_FTPHostCvar.GetString(g_FTPHost, sizeof(g_FTPHost));
    g_FTPPort = g_FTPPortCvar.IntValue;
    g_FTPUsernameCvar.GetString(g_FTPUsername, sizeof(g_FTPUsername));
    g_FTPPasswordCvar.GetString(g_FTPPassword, sizeof(g_FTPPassword));
    
    Format(zippedFile, sizeof(zippedFile), "%s", filename);
    Format(remoteDemoPath, sizeof(remoteDemoPath), "%s/%s", g_FTPHost, zippedFile);
    LogDebug("Our File is: %s and remote demo path of %s", zippedFile, remoteDemoPath);
    System2FTPRequest ftpRequest = new System2FTPRequest(FtpResponseCallback, remoteDemoPath);
    ftpRequest.AppendToFile = false;
    ftpRequest.CreateMissingDirs = true;
    ftpRequest.SetAuthentication(g_FTPUsername, g_FTPPassword);
    ftpRequest.SetPort(g_FTPPort);
    ftpRequest.SetProgressCallback(FtpProgressCallback);
    LogDebug("Our File is: %s", zippedFile);

    ftpRequest.SetInputFile(zippedFile);
    ftpRequest.StartRequest(); 
  } else{
    LogDebug("FTP Uploads Disabled OR Filename was empty (no demo to upload). Change config to enable.");
  }
}


public void FtpProgressCallback(System2FTPRequest request, int dlTotal, int dlNow, int ulTotal, int ulNow) {
  char file[PLATFORM_MAX_PATH];
  request.GetInputFile(file, sizeof(file));
  if (strlen(file) > 0) {
      LogDebug("Uploading %s file with %d bytes total, %d now", file, ulTotal, ulNow);
  }
}  

public void FtpResponseCallback(bool success, const char[] error, System2FTPRequest request, System2FTPResponse response) {
    if (success || StrContains(error, "Uploaded unaligned file size") > -1) {
        char file[PLATFORM_MAX_PATH];
        request.GetInputFile(file, sizeof(file));
        if (strlen(file) > 0) {
            if (DeleteFileIfExists(file)) {
                LogDebug("Deleted file after complete.");
            }
        }
    } else{
      LogError("There was a problem: %s", error);
    }
}

public void Get5_OnMapPicked(MatchTeam team, const char[] map){
  char teamString[64];
  GetTeamString(team, teamString, sizeof(teamString));
  LogDebug("Map Pick START team %s map picked %s", team, map);
  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/vetoUpdate", g_MatchID);
  if (req != INVALID_HANDLE) {
      AddStringParam(req, "map", map);
      AddStringParam(req, "teamString", teamString);
      AddStringParam(req, "pick_or_veto", "pick");
      SteamWorks_SendHTTPRequest(req);
  }
  LogDebug("Accepted Map Pick.");
}

public void Get5_OnSeriesResult(MatchTeam seriesWinner, int team1MapScore, int team2MapScore) {
  char winnerString[64];
  GetTeamString(seriesWinner, winnerString, sizeof(winnerString));

  KeyValues kv = new KeyValues("Stats");
  Get5_GetMatchStats(kv);
  bool forfeit = kv.GetNum(STAT_SERIES_FORFEIT, 0) != 0;
  delete kv;

  Handle req = CreateRequest(k_EHTTPMethodPOST, "match/%d/finish", g_MatchID);
  if (req != INVALID_HANDLE) {
    AddStringParam(req, "winner", winnerString);
    AddIntParam(req, "forfeit", forfeit);
    SteamWorks_SendHTTPRequest(req);
  }

  g_APIKeyCvar.SetString("");
}

public void Get5_OnRoundStatsUpdated() {
  if (Get5_GetGameState() == Get5State_Live) {
    UpdateRoundStats(MapNumber());
  }
}

static int MapNumber() {
  int t1, t2, n;
  int buf;
  Get5_GetTeamScores(MatchTeam_Team1, t1, buf);
  Get5_GetTeamScores(MatchTeam_Team2, t2, buf);
  Get5_GetTeamScores(MatchTeam_TeamNone, n, buf);
  return t1 + t2 + n;
}
