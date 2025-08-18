<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import AppCard from './AppCard.svelte';
  import type { AppInfo } from '../types/app';
  
  export let favorites: AppInfo[] = [];
</script>

<section class="favorites-section">
  <div class="favorites-header">
    <h2>⭐ Favorites ({favorites.length})</h2>
  </div>
  
  {#if favorites.length > 0}
    <div class="favorites-grid">
      {#each favorites as app (app.name)}
        <AppCard {app} />
      {/each}
    </div>
  {:else}
    <div class="no-favorites">
      <p>No favorite applications yet.</p>
      <p>Click the star icon on any app to add it to your favorites.</p>
    </div>
  {/if}
</section>

<style>
  .favorites-section {
    background: #161b22;
    border: 1px solid #30363d;
    border-radius: 6px;
    padding: 24px;
    margin-bottom: 24px;
  }
  
  .favorites-header {
    margin-bottom: 20px;
  }
  
  .favorites-header h2 {
    margin: 0;
    color: #f0f6fc;
    font-size: 1.25rem;
    font-weight: 600;
    display: flex;
    align-items: center;
    gap: 8px;
  }
  
  .favorites-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 16px;
  }
  
  .no-favorites {
    text-align: center;
    padding: 40px;
    color: #8b949e;
  }
  
  .no-favorites p {
    margin: 8px 0;
  }
  
  @media (max-width: 768px) {
    .favorites-section {
      padding: 20px;
    }
    
    .favorites-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
