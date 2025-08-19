<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher<{
    search: string;
  }>();
  
  let searchQuery = '';
  let searchTimeout: number;
  
  function handleInput() {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
      dispatch('search', searchQuery);
    }, 150);
  }
  
  function clearSearch() {
    searchQuery = '';
    dispatch('search', '');
  }
</script>

<div>
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
    <circle cx="11" cy="11" r="8"/>
    <path d="m21 21-4.35-4.35"/>
  </svg>
  
  <input
    type="text"
    placeholder="Search applications, commands, or executables..."
    bind:value={searchQuery}
    on:input={handleInput}
    aria-label="Search applications"
  />
  
  {#if searchQuery}
    <button
      type="button"
      on:click={clearSearch}
      aria-label="Clear search"
    >
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <line x1="18" y1="6" x2="6" y2="18"/>
        <line x1="6" y1="6" x2="18" y2="18"/>
      </svg>
    </button>
  {/if}
</div>


