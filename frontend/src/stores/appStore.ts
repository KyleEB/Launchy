import { writable, derived } from 'svelte/store';
import { appService } from '../services/AppService';
import type { AppInfo } from '../types/app';

// Create writable stores for reactive state
export const apps = writable<AppInfo[]>([]);
export const searchQuery = writable<string>('');
export const isLoading = writable<boolean>(true);
export const error = writable<string>('');
export const loadingState = writable<"initializing" | "loading" | "loaded" | "error">("initializing");

// Derived store for filtered apps
export const filteredApps = derived(
  [apps, searchQuery],
  ([$apps, $searchQuery]) => {
    return appService.searchApps($searchQuery);
  }
);

// Store actions that update both the service and the stores
export const appStore = {
  async loadApps(): Promise<void> {
    await appService.loadApps();
    
    // Update stores with service state
    apps.set(appService.getApps());
    isLoading.set(appService.isLoading());
    error.set(appService.getError());
    loadingState.set(appService.getLoadingState());
  },

  setSearchQuery(query: string): void {
    appService.setSearchQuery(query);
    searchQuery.set(query);
  },

  async launchApp(execPath: string): Promise<void> {
    await appService.launchApp(execPath);
  },

  toggleFavorite(appName: string): void {
    appService.toggleFavorite(appName);
  },

  // Getters that return store values
  getApps() {
    return appService.getApps();
  },

  getFilteredApps() {
    return appService.getFilteredApps();
  },

  getSearchQuery() {
    return appService.getSearchQuery();
  },

  getIsLoading() {
    return appService.isLoading();
  },

  getError() {
    return appService.getError();
  },

  getLoadingState() {
    return appService.getLoadingState();
  }
};
