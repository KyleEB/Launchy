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
>
  <h3>{app.name}</h3>
  <button
    type="button"
    on:click={handleToggleFavorite}
    aria-label="{app.isFavorite ? 'Remove from favorites' : 'Add to favorites'}"
  >
    <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
      <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
    </svg>
  </button>
  
  {#if app.comment}
    <p>{app.comment}</p>
  {/if}
  
  {#if app.categories}
    {#each app.categories.split(';').slice(0, 2) as category}
      {#if category.trim()}
        <span>{category.trim()}</span>
      {/if}
    {/each}
  {/if}
  
  <p>{app.exec}</p>
</div>


