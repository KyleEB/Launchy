import type { AppInfo } from '../types/app';

export interface AppServiceInterface {
  loadApps(): Promise<void>;
  searchApps(query: string): AppInfo[];
  toggleFavorite(appName: string): void;
  launchApp(execPath: string): Promise<void>;
  getApps(): AppInfo[];
  getFilteredApps(): AppInfo[];
  getSearchQuery(): string;
  setSearchQuery(query: string): void;
  isLoading(): boolean;
  getError(): string;
  getLoadingState(): "initializing" | "loading" | "loaded" | "error";
}

export class AppService implements AppServiceInterface {
  private apps: AppInfo[] = [];
  private searchQuery: string = "";
  private _isLoading: boolean = true;
  private error: string = "";
  private loadingState: "initializing" | "loading" | "loaded" | "error" = "initializing";
  private bindings: any = null;

  async loadApps(): Promise<void> {
    try {
      this.loadingState = "loading";
      this._isLoading = true;
      this.error = "";
      
      // Import bindings
      const bindingsModule = await import("../../bindings/changeme/backend/service/index.js");
      this.bindings = bindingsModule.AppService;
      
      // Get apps with timeout
      const appsResult = await Promise.race([
        this.bindings.GetApps(),
        new Promise((_, reject) => 
          setTimeout(() => reject(new Error("GetApps timeout after 10 seconds")), 10000)
        )
      ]);
      
      // Parse the result - it might be a JSON string or already parsed
      let parsedApps: AppInfo[];
      if (typeof appsResult === 'string') {
        try {
          parsedApps = JSON.parse(appsResult);
        } catch (parseErr) {
          console.error("Failed to parse apps result:", parseErr);
          throw new Error("Invalid JSON response from backend");
        }
      } else if (Array.isArray(appsResult)) {
        parsedApps = appsResult;
      } else {
        console.error("Unexpected apps result format:", appsResult);
        throw new Error("Unexpected response format from backend");
      }
      
      this.apps = parsedApps;
      this.loadingState = "loaded";
      this._isLoading = false;
    } catch (err) {
      console.error("Failed to load apps:", err);
      const errorMsg = `Failed to load applications: ${err.message}`;
      this.error = errorMsg;
      this.loadingState = "error";
      this._isLoading = false;
    }
  }

  searchApps(query: string): AppInfo[] {
    if (query.trim() === "") {
      return this.apps;
    }
    
    const searchQuery = query.toLowerCase();
    return this.apps.filter(app => 
      app.name.toLowerCase().includes(searchQuery) ||
      (app.comment && app.comment.toLowerCase().includes(searchQuery)) ||
      (app.categories && app.categories.toLowerCase().includes(searchQuery)) ||
      (app.exec && app.exec.toLowerCase().includes(searchQuery))
    );
  }

  toggleFavorite(appName: string): void {
    // TODO: Implement favorite functionality
    // This could involve updating local storage, making API calls, etc.
    console.log("Toggle favorite for:", appName);
  }

  async launchApp(execPath: string): Promise<void> {
    try {
      if (this.bindings && this.bindings.LaunchApp) {
        await this.bindings.LaunchApp(execPath);
      } else {
        console.error("Bindings not available for launching app");
      }
    } catch (err) {
      console.error("Failed to launch app:", err);
    }
  }

  getApps(): AppInfo[] {
    return this.apps;
  }

  getFilteredApps(): AppInfo[] {
    return this.searchApps(this.searchQuery);
  }

  getSearchQuery(): string {
    return this.searchQuery;
  }

  setSearchQuery(query: string): void {
    this.searchQuery = query;
  }

  isLoading(): boolean {
    return this._isLoading;
  }

  getError(): string {
    return this.error;
  }

  getLoadingState(): "initializing" | "loading" | "loaded" | "error" {
    return this.loadingState;
  }
}

// Export a singleton instance
export const appService = new AppService();
