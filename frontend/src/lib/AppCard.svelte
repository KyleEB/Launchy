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

<div class="app-card" on:click={handleLaunch} on:keydown={handleKeydown} role="button" tabindex="0">
  <div class="app-info">
    <h3 class="app-name">{app.name}</h3>
    {#if app.comment}
      <p class="app-comment">{app.comment}</p>
    {/if}
    {#if app.categories}
      <div class="app-categories">
        {#each app.categories.split(';').slice(0, 2) as category}
          {#if category.trim()}
            <span class="category-tag">{category.trim()}</span>
          {/if}
        {/each}
      </div>
    {/if}
  </div>
  
  <button
    type="button"
    class="favorite-button"
    class:favorited={app.isFavorite}
    on:click={handleToggleFavorite}
    aria-label="{app.isFavorite ? 'Remove from favorites' : 'Add to favorites'}"
  >
    <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
      <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
    </svg>
  </button>
</div>

<style>
  .app-card {
    display: flex;
    align-items: center;
    padding: 16px;
    background: #21262d;
    border: 1px solid #30363d;
    border-radius: 6px;
    cursor: pointer;
    transition: all 0.2s ease;
    position: relative;
  }
  
  .app-card:hover {
    border-color: #58a6ff;
    background: #161b22;
  }
  
  .app-card:focus {
    outline: none;
    border-color: #58a6ff;
    box-shadow: 0 0 0 3px rgba(56, 139, 253, 0.15);
  }
  

  
  .app-info {
    flex: 1;
    min-width: 0;
  }
  
  .app-name {
    margin: 0 0 4px 0;
    font-size: 14px;
    font-weight: 600;
    color: #f0f6fc;
    line-height: 1.2;
  }
  
  .app-comment {
    margin: 0 0 8px 0;
    font-size: 12px;
    color: #8b949e;
    line-height: 1.3;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  
  .app-categories {
    display: flex;
    gap: 4px;
    flex-wrap: wrap;
  }
  
  .category-tag {
    font-size: 11px;
    padding: 2px 6px;
    background: #30363d;
    color: #8b949e;
    border-radius: 3px;
    white-space: nowrap;
  }
  
  .favorite-button {
    background: none;
    border: none;
    color: #8b949e;
    cursor: pointer;
    padding: 4px;
    border-radius: 3px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    flex-shrink: 0;
    margin-left: 8px;
  }
  
  .favorite-button:hover {
    background: #21262d;
    color: #fbb040;
  }
  
  .favorite-button.favorited {
    color: #fbb040;
  }
  
  .favorite-button:focus {
    outline: 2px solid #58a6ff;
    outline-offset: 2px;
  }
  
  @media (max-width: 768px) {
    .app-card {
      padding: 12px;
    }
    
    .app-name {
      font-size: 13px;
    }
    
    .app-comment {
      font-size: 11px;
    }
  }
</style>
