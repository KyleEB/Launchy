<script lang="ts">
  import SearchBar from "./components/SearchBar.svelte";
  import AppGrid from "./components/AppGrid.svelte";
  import ErrorMessage from "./components/ErrorMessage.svelte";
  import { Rocket } from "lucide-svelte";
  import { appStore, filteredApps, isLoading, error } from "./stores/appStore";
  import { onMount } from 'svelte';

  let selectedIndex = 0;

  // Search functionality
  function handleSearch(event: CustomEvent<string>) {
    appStore.setSearchQuery(event.detail);
    selectedIndex = 0; // Reset selection when search changes
  }

  // Handle launch event from AppCard component
  function handleLaunch(event: CustomEvent<string>) {
    appStore.launchApp(event.detail);
  }



  // Handle keyboard navigation
  function handleKeydown(event: KeyboardEvent) {
    if ($filteredApps.length === 0) return;

    switch (event.key) {
      case 'ArrowDown':
        event.preventDefault();
        selectedIndex = Math.min(selectedIndex + 1, $filteredApps.length - 1);
        break;
      case 'ArrowUp':
        event.preventDefault();
        selectedIndex = Math.max(selectedIndex - 1, 0);
        break;
      case 'Enter':
        event.preventDefault();
        if ($filteredApps[selectedIndex]) {
          appStore.launchApp($filteredApps[selectedIndex].exec);
        }
        break;
    }
  }

  // Load apps when component mounts
  onMount(() => {
    appStore.loadApps();
  });
</script>

<main class="min-h-screen bg-[#0d1117] text-[#c9d1d9]" on:keydown={handleKeydown} tabindex="0">
    <header class="bg-[#161b22] border-b border-[#30363d] px-4 py-2 shadow-sm">
    <div class="max-w-6xl mx-auto flex items-center justify-center space-x-6">
      <div class="flex items-center space-x-2">
        <Rocket class="text-[#58a6ff] w-6 h-6" />
        <h1 class="text-lg font-bold">Launchy</h1>
      </div>
      <div class="w-100">
        <SearchBar on:search={handleSearch} />
      </div>
      <div class="flex items-center space-x-2 text-sm text-gray-400 bg-gray-800 px-3 py-1 rounded-md border border-gray-600">
        <span>↑↓ Navigate</span>
      </div>
    </div>
  </header>

  <div class="max-w-7xl mx-auto px-4 py-6">
    {#if $error}
      <ErrorMessage error={$error} onRetry={() => appStore.loadApps()} />
    {:else}

    {#if !$isLoading}
      <AppGrid 
        apps={$filteredApps}
        {selectedIndex}
        on:launch={handleLaunch}
      />
    {/if}
  {/if}
    </div>
</main>


