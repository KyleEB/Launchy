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

<main class="min-h-screen bg-[#0d1117] text-[#c9d1d9]">
    <header class="bg-[#161b22] border-b border-[#30363d] px-6 py-4 shadow-sm">
    <div class="max-w-6xl mx-auto flex items-center justify-between">
      <div class="flex items-center space-x-3">
        <Rocket class="text-[#58a6ff] p-10 w-40 h-40" />
        <h1 class="text-2xl font-bold">Launchy</h1>
      </div>
      <div class="flex-1 max-w-md ml-8">
        <SearchBar on:search={handleSearch} />
      </div>
    </div>
  </header>

  <div class="max-w-7xl mx-auto px-4 py-6">
    {#if $error}
      <ErrorMessage error={$error} onRetry={() => appStore.loadApps()} />
    {:else}

    {#if $filteredApps.length === 0 && !$isLoading}
      <div class="text-center py-16">
        <svg class="w-12 h-12 mx-auto text-[#484f58] mb-4" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
        </svg>
        <p class="text-lg font-medium text-[#f0f6fc] mb-2">No applications found</p>
        {#if appStore.getSearchQuery()}
          <p class="text-[#8b949e]">Try adjusting your search terms</p>
        {/if}
      </div>
    {:else if $filteredApps.length > 0}
      <div class="grid grid-cols-2 max-w-2xl mx-auto p-4" style="gap: 24px;">
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


