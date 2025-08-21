// Application information interface
export interface AppInfo {
  name: string;
  exec: string;
  icon: string;
  comment: string;
  categories: string;
  isFavorite: boolean;
}

// Search event interface
export interface SearchEvent {
  query: string;
}

// Launch event interface
export interface LaunchEvent {
  execPath: string;
}

// Toggle favorite event interface
export interface ToggleFavoriteEvent {
  appName: string;
}

// Application service interface
export interface AppService {
  GetApps(): Promise<AppInfo[]>;
  GetFavorites(): Promise<AppInfo[]>;
  SearchApps(query: string): Promise<AppInfo[]>;
  LaunchApp(execPath: string): Promise<void>;
  ToggleFavorite(appName: string): Promise<void>;
}
