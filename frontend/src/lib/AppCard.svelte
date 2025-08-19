<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { AppInfo } from '../types/app';
  
  export let app: AppInfo;
  
  const dispatch = createEventDispatcher<{
    launch: string;
    toggleFavorite: string;
  }>();
  
  function handleLaunch() {
    dispatch('launch', app.exec);
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      handleLaunch();
    }
  }
  
  function handleToggleFavorite(event: Event) {
    event.stopPropagation();
    dispatch('toggleFavorite', app.name);
  }
</script>

<div 
  on:click={handleLaunch} 
  on:keydown={handleKeydown} 
  role="button" 
  tabindex="0"
  class="bg-[#161b22] border border-[#30363d] hover:border-[#58a6ff] hover:bg-[#1c2128] transition-all duration-200 cursor-pointer focus:outline-none focus:ring-2 focus:ring-[#58a6ff] focus:ring-opacity-50"
  style="border-radius: 16px; padding: 12px;"
>
  <div class="flex items-start justify-between mb-1">
    <h3 class="text-sm font-medium text-[#f0f6fc] truncate">{app.name}</h3>
    <button
      type="button"
      on:click={handleToggleFavorite}
      aria-label="{app.isFavorite ? 'Remove from favorites' : 'Add to favorites'}"
      class="flex-shrink-0 ml-1 p-0.5 text-[#8b949e] hover:text-[#f78166] transition-colors duration-200"
    >
      <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor" class="{app.isFavorite ? 'text-[#f78166]' : ''}">
        <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
      </svg>
    </button>
  </div>
  
  {#if app.comment}
    <p class="text-xs text-[#8b949e] mb-1 line-clamp-1">{app.comment}</p>
  {/if}
  
  {#if app.categories}
    <div class="flex flex-wrap gap-1 mb-1">
      {#each app.categories.split(';').slice(0, 2) as category}
        {#if category.trim()}
          <span class="inline-block px-1.5 py-0.5 text-xs font-medium bg-[#21262d] text-[#c9d1d9] rounded border border-[#30363d]">{category.trim()}</span>
        {/if}
      {/each}
    </div>
  {/if}
  
  <p class="text-xs text-[#a855f7] font-mono truncate">{app.exec}</p>
</div>


