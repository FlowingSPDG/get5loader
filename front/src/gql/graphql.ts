/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  DateTime: { input: any; output: any; }
  SteamID: { input: any; output: any; }
};

export type GameServer = {
  __typename?: 'GameServer';
  Ip: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  port: Scalars['Int']['output'];
  public: Scalars['Boolean']['output'];
};

export type LoginToken = {
  __typename?: 'LoginToken';
  token: Scalars['String']['output'];
};

export type MapStats = {
  __typename?: 'MapStats';
  endedAt?: Maybe<Scalars['DateTime']['output']>;
  id: Scalars['ID']['output'];
  mapName: Scalars['String']['output'];
  mapNumber: Scalars['Int']['output'];
  matchId: Scalars['ID']['output'];
  playerstats: Array<PlayerStats>;
  startedAt?: Maybe<Scalars['DateTime']['output']>;
  team1score: Scalars['Int']['output'];
  team2score: Scalars['Int']['output'];
  winner?: Maybe<Scalars['ID']['output']>;
};

export type Match = {
  __typename?: 'Match';
  endedAt?: Maybe<Scalars['DateTime']['output']>;
  forfeit?: Maybe<Scalars['Boolean']['output']>;
  id: Scalars['ID']['output'];
  mapStats: Array<MapStats>;
  maxMaps: Scalars['Int']['output'];
  skipVeto: Scalars['Boolean']['output'];
  startedAt?: Maybe<Scalars['DateTime']['output']>;
  team1: Team;
  team1Score: Scalars['Int']['output'];
  team2: Team;
  team2Score: Scalars['Int']['output'];
  title: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
  winner: Scalars['ID']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addServer: GameServer;
  createMatch: Match;
  registerTeam: Team;
};


export type MutationAddServerArgs = {
  input: NewGameServer;
};


export type MutationCreateMatchArgs = {
  input: NewMatch;
};


export type MutationRegisterTeamArgs = {
  input: NewTeam;
};

export type NewGameServer = {
  Ip: Scalars['String']['input'];
  RconPassword: Scalars['String']['input'];
  name: Scalars['String']['input'];
  port: Scalars['Int']['input'];
  public: Scalars['Boolean']['input'];
};

export type NewMatch = {
  maxMaps: Scalars['Int']['input'];
  serverID: Scalars['ID']['input'];
  skipVeto: Scalars['Boolean']['input'];
  team1: Scalars['ID']['input'];
  team2: Scalars['ID']['input'];
  title: Scalars['String']['input'];
};

export type NewPlayer = {
  name: Scalars['String']['input'];
  steamId: Scalars['SteamID']['input'];
  teamid: Scalars['ID']['input'];
};

export type NewPlayerForTeam = {
  name: Scalars['String']['input'];
  steamId: Scalars['SteamID']['input'];
};

export type NewTeam = {
  flag: Scalars['String']['input'];
  logo: Scalars['String']['input'];
  name: Scalars['String']['input'];
  players?: InputMaybe<Array<NewPlayerForTeam>>;
  public: Scalars['Boolean']['input'];
  tag: Scalars['String']['input'];
};

export type NewUser = {
  admin: Scalars['Boolean']['input'];
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
  steamId: Scalars['SteamID']['input'];
};

export type Player = {
  __typename?: 'Player';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  steamId: Scalars['SteamID']['output'];
  teamId: Scalars['ID']['output'];
};

export type PlayerStats = {
  __typename?: 'PlayerStats';
  assists: Scalars['Int']['output'];
  bombDefuses: Scalars['Int']['output'];
  bombPlants: Scalars['Int']['output'];
  damage: Scalars['Int']['output'];
  deaths: Scalars['Int']['output'];
  firstDeathCT: Scalars['Int']['output'];
  firstDeathT: Scalars['Int']['output'];
  firstKillCT: Scalars['Int']['output'];
  firstKillT: Scalars['Int']['output'];
  flashBangAssists: Scalars['Int']['output'];
  headshotKills: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  k1: Scalars['Int']['output'];
  k2: Scalars['Int']['output'];
  k3: Scalars['Int']['output'];
  k4: Scalars['Int']['output'];
  k5: Scalars['Int']['output'];
  kills: Scalars['Int']['output'];
  mapstatsId: Scalars['ID']['output'];
  matchId: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  roundsPlayed: Scalars['Int']['output'];
  steamId: Scalars['SteamID']['output'];
  suicides: Scalars['Int']['output'];
  v1: Scalars['Int']['output'];
  v2: Scalars['Int']['output'];
  v3: Scalars['Int']['output'];
  v4: Scalars['Int']['output'];
  v5: Scalars['Int']['output'];
};

export type Query = {
  __typename?: 'Query';
  getMatch: Match;
  getMatchesByMe: Array<Match>;
  getMatchesByUser: Array<Match>;
  getMe: User;
  getPublicServers: Array<GameServer>;
  getServer: GameServer;
  getTeam: Team;
  getTeamsByUser: Array<Team>;
  getUser: User;
};


export type QueryGetMatchArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetMatchesByUserArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetServerArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetTeamArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetUserArgs = {
  id: Scalars['ID']['input'];
};

export type Team = {
  __typename?: 'Team';
  flag: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  logo: Scalars['String']['output'];
  name: Scalars['String']['output'];
  players: Array<Player>;
  public: Scalars['Boolean']['output'];
  tag: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type User = {
  __typename?: 'User';
  admin: Scalars['Boolean']['output'];
  gameservers: Array<GameServer>;
  id: Scalars['ID']['output'];
  matches: Array<Match>;
  name: Scalars['String']['output'];
  steamId: Scalars['SteamID']['output'];
  teams: Array<Team>;
};

export type UserLoginId = {
  ID: Scalars['ID']['input'];
  password: Scalars['String']['input'];
};

export type UserLoginSteamId = {
  password: Scalars['String']['input'];
  steamId: Scalars['SteamID']['input'];
};
