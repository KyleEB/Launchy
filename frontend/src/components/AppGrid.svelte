<script lang="ts">
  import AppCard from "./AppCard.svelte";
  import type { AppInfo } from '../types/app';

  const { apps, selectedIndex, onLaunch } = $props<{
    apps: AppInfo[];
    selectedIndex: number;
    onLaunch: (exec: string) => void;
  }>();

  // Handle launch event from AppCard component
  function handleLaunch(exec: string) {
    onLaunch(exec);
  }


</script>

{#if apps.length === 0}
  <div class="text-center py-16">
    <svg class="w-12 h-12 mx-auto text-[#484f58] mb-4" fill="currentColor" viewBox="0 0 24 24">
      <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
    </svg>
    <p class="text-lg font-medium text-[#f0f6fc] mb-2">No applications found</p>
  </div>
{:else}
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 max-w-6xl mx-auto p-4" style="gap: 24px;">
      {#each apps as app, index (app.exec + '-' + index)}
        <AppCard 
          {app}
          isSelected={index === selectedIndex}
          onLaunch={handleLaunch}
        />
      {/each}
    </div>
{/if}
