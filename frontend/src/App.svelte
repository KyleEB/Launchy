<script lang="ts">
  import SearchBar from "./lib/SearchBar.svelte";
  import AppCard from "./lib/AppCard.svelte";
  import ErrorMessage from "./lib/ErrorMessage.svelte";
  import { Rocket } from "lucide-svelte";
  import { appStore, filteredApps, isLoading, error } from "./stores/appStore";
  import { onMount } from 'svelte';

  // Search functionality
  function handleSearch(event: CustomEvent<string>) {
    appStore.setSearchQuery(event.detail);
  }

  // Handle launch event from AppCard component
  function handleLaunch(event: CustomEvent<string>) {
    appStore.launchApp(event.detail);
  }

  // Handle toggle favorite event from AppCard component
  function handleToggleFavorite(event: CustomEvent<string>) {
    appStore.toggleFavorite(event.detail);
  }

  // Load apps when component mounts
  onMount(() => {
    appStore.loadApps();
  });
</script>

<main>
  <header>
    <Rocket />
    <h1>Launchy</h1>
    <p>Application Launcher for CachyOS</p>
  </header>

  <div>
    {#if $error}
      <ErrorMessage error={$error} onRetry={() => appStore.loadApps()} />
    {:else}
      <SearchBar on:search={handleSearch} />

    {#if $filteredApps.length === 0 && !$isLoading}
      <div>
        <svg fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
        </svg>
        <p>No applications found</p>
        {#if appStore.getSearchQuery()}
          <p>Try adjusting your search terms</p>
        {/if}
      </div>
    {:else if $filteredApps.length > 0}
      <div style="gap: 20px;">
        {#each $filteredApps as app, index (app.exec + '-' + index)}
          <AppCard 
            {app}
            on:launch={handleLaunch}
            on:toggleFavorite={handleToggleFavorite}
          />
        {/each}
      </div>
    {:else}
      <!-- Loading: intentionally render nothing -->
    {/if}
  {/if}
    </div>
</main>


