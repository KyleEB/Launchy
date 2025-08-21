<script lang="ts">
  import SearchBar from "./components/SearchBar.svelte";
  import AppGrid from "./components/AppGrid.svelte";
  import ErrorMessage from "./components/ErrorMessage.svelte";
  import { Rocket } from "lucide-svelte";
  import { appStore, filteredApps, isLoading, error } from "./stores/appStore";
  import { onMount } from 'svelte';

  let selectedIndex = $state(0);
  let searchBarComponent = $state<any>(null);

  // Search functionality
  function handleSearch(query: string) {
    appStore.setSearchQuery(query);
    selectedIndex = 0; // Reset selection when search changes
  }

  // Handle launch event from AppCard component
  function handleLaunch(exec: string) {
    appStore.launchApp(exec);
  }

  // Handle keyboard navigation and text input
  function handleKeydown(event: KeyboardEvent) {
    // Check if the target is already an input element
    const target = event.target as HTMLElement;
    if (target.tagName === 'INPUT' || target.tagName === 'TEXTAREA') {
      return; // Let input elements handle their own events
    }

    // Handle navigation keys
    if (event.key === 'ArrowDown') {
      event.preventDefault();
      if ($filteredApps.length > 0) {
        selectedIndex = Math.min(selectedIndex + 1, $filteredApps.length - 1);
      }
      return;
    }

    if (event.key === 'ArrowUp') {
      event.preventDefault();
      if ($filteredApps.length > 0) {
        selectedIndex = Math.max(selectedIndex - 1, 0);
      }
      return;
    }

    if (event.key === 'Enter') {
      event.preventDefault();
      if ($filteredApps[selectedIndex]) {
        appStore.launchApp($filteredApps[selectedIndex].exec);
      }
      return;
    }

    // Handle text input - send to search bar
    if (event.key.length === 1 && !event.ctrlKey && !event.metaKey && !event.altKey) {
      event.preventDefault();
      if (searchBarComponent) {
        searchBarComponent.focusAndType(event.key);
      }
    }
  }

  // Load apps when component mounts
  onMount(() => {
    appStore.loadApps();
  });
</script>

<div class="min-h-screen bg-[#0d1117] text-[#c9d1d9]" onkeydown={handleKeydown} contenteditable="true" role="application">
    <header class="bg-[#161b22] border-b border-[#30363d] px-4 py-2 shadow-sm">
    <div class="max-w-6xl mx-auto flex items-center justify-center space-x-6">
      <div class="flex items-center space-x-2">
        <Rocket class="text-[#58a6ff] w-6 h-6" />
        <h1 class="text-lg font-bold">Launchy</h1>
      </div>
      <div class="w-100">
        <SearchBar bind:this={searchBarComponent} onSearch={handleSearch} onEnter={() => {
          if ($filteredApps[selectedIndex]) {
            appStore.launchApp($filteredApps[selectedIndex].exec);
          }
        }} onArrow={(direction) => {
          if (direction === 'ArrowDown') {
            selectedIndex = Math.min(selectedIndex + 1, $filteredApps.length - 1);
          } else if (direction === 'ArrowUp') {
            selectedIndex = Math.max(selectedIndex - 1, 0);
          }
        }} />
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
        onLaunch={handleLaunch}
      />
    {/if}
  {/if}
    </div>
</div>


