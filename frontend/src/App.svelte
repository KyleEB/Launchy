<script lang="ts">
  import AppCard from "./lib/AppCard.svelte";
  import SearchBar from "./lib/SearchBar.svelte";
  import FavoritesSection from "./lib/FavoritesSection.svelte";
  import type { AppInfo } from "./types/app";

  // Dynamic import to avoid TypeScript issues
  let AppService: any = null;
  
  let apps: AppInfo[] = [];
  let favorites: AppInfo[] = [];
  let filteredApps: AppInfo[] = [];
  let searchQuery = "";
  let isLoading = true;
  let error = "";

  // Load bindings and apps on component mount
  async function loadApps() {
    try {
      // Dynamic import of bindings
      const bindings = await import("../bindings/changeme/index.js");
      AppService = bindings.AppService;
      
      const [appsResult, favoritesResult] = await Promise.all([
        AppService.GetApps(),
        AppService.GetFavorites()
      ]);
      
      apps = appsResult;
      favorites = favoritesResult;
      filteredApps = apps;
      isLoading = false;
    } catch (err) {
      console.error("Failed to load apps:", err);
      error = "Failed to load applications. Please try again.";
      isLoading = false;
    }
  }

  // Search functionality
  async function handleSearch(event: CustomEvent<string>) {
    const query = event.detail;
    searchQuery = query;
    if (query.trim() === "") {
      filteredApps = apps;
    } else {
      try {
        filteredApps = await AppService.SearchApps(query);
      } catch (err) {
        console.error("Failed to search apps:", err);
        filteredApps = [];
      }
    }
  }

  // Launch app
  async function launchApp(event: CustomEvent<string>) {
    const execPath = event.detail;
    try {
      await AppService.LaunchApp(execPath);
    } catch (err) {
      console.error("Failed to launch app:", err);
    }
  }

  // Toggle favorite
  async function toggleFavorite(event: CustomEvent<string>) {
    const appName = event.detail;
    try {
      await AppService.ToggleFavorite(appName);
      // Reload apps and favorites to get updated state
      const [appsResult, favoritesResult] = await Promise.all([
        AppService.GetApps(),
        AppService.GetFavorites()
      ]);
      apps = appsResult;
      favorites = favoritesResult;
      // Update filtered apps if we're currently showing all apps
      if (searchQuery.trim() === "") {
        filteredApps = apps;
      }
    } catch (err) {
      console.error("Failed to toggle favorite:", err);
    }
  }

  // Load apps when component mounts
  loadApps();
</script>

<main class="app">
  <div class="header">
    <h1>🚀 Launchy</h1>
    <p>Application Launcher for CachyOS</p>
  </div>

  {#if error}
    <div class="error-message">
      <p>{error}</p>
      <button on:click={loadApps}>Retry</button>
    </div>
  {:else}
    <SearchBar on:search={handleSearch} />

    {#if favorites.length > 0}
      <FavoritesSection 
        {favorites} 
        on:launch={launchApp}
        on:toggleFavorite={toggleFavorite}
      />
    {/if}

    <div class="apps-section">
      <h2>Applications ({filteredApps.length})</h2>
      
      {#if isLoading}
        <div class="loading">
          <div class="spinner"></div>
          <p>Loading applications...</p>
        </div>
      {:else if filteredApps.length === 0}
        <div class="no-results">
          <p>No applications found.</p>
          {#if searchQuery}
            <p>Try adjusting your search terms.</p>
          {/if}
        </div>
      {:else}
        <div class="apps-grid">
          {#each filteredApps as app (app.name)}
            <AppCard 
              {app} 
              on:launch={launchApp}
              on:toggleFavorite={toggleFavorite}
            />
          {/each}
        </div>
      {/if}
    </div>
  {/if}
</main>

<style>
  .app {
    min-height: 100vh;
    background-color: #0d1117;
    color: #c9d1d9;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans', Helvetica, Arial, sans-serif;
    padding: 24px;
  }

  .header {
    text-align: center;
    margin-bottom: 32px;
    padding-bottom: 24px;
    border-bottom: 1px solid #30363d;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .header h1 {
    font-size: 2.5rem;
    margin: 0 0 8px 0;
    font-weight: 600;
    color: #f0f6fc;
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .header p {
    font-size: 1.1rem;
    margin: 0;
    color: #8b949e;
  }

  .error-message {
    background: #21262d;
    border: 1px solid #f85149;
    border-radius: 6px;
    padding: 16px;
    margin-bottom: 24px;
    text-align: center;
    color: #f85149;
  }

  .error-message button {
    background: #f85149;
    color: #ffffff;
    border: none;
    padding: 8px 16px;
    border-radius: 6px;
    cursor: pointer;
    margin-top: 12px;
    font-size: 14px;
    font-weight: 500;
  }

  .error-message button:hover {
    background: #da3633;
  }

  .apps-section {
    background: #161b22;
    border: 1px solid #30363d;
    border-radius: 6px;
    padding: 24px;
    margin-top: 24px;
  }

  .apps-section h2 {
    margin: 0 0 20px 0;
    color: #f0f6fc;
    font-size: 1.25rem;
    font-weight: 600;
  }

  .apps-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 16px;
  }

  .loading {
    text-align: center;
    padding: 40px;
    color: #8b949e;
  }

  .spinner {
    width: 32px;
    height: 32px;
    border: 3px solid #30363d;
    border-top: 3px solid #58a6ff;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto 16px;
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }

  .no-results {
    text-align: center;
    padding: 40px;
    color: #8b949e;
  }

  .no-results p {
    margin: 8px 0;
  }

  @media (max-width: 768px) {
    .app {
      padding: 16px;
    }

    .header h1 {
      font-size: 2rem;
    }

    .apps-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
