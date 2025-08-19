<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import AppCard from './AppCard.svelte';
  import type { AppInfo } from '../types/app';
  
  export let favorites: AppInfo[] = [];
</script>

<section>
  <h2>⭐ Favorites ({favorites.length})</h2>
  
  {#if favorites.length > 0}
    {#each favorites as app (app.name)}
      <AppCard {app} />
    {/each}
  {:else}
    <p>No favorite applications yet.</p>
    <p>Click the star icon on any app to add it to your favorites.</p>
  {/if}
</section>


