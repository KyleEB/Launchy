<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher<{
    search: string;
  }>();
  
  let searchQuery = '';
  let searchTimeout: number;
  
  function handleInput(event: Event) {
    const target = event.target as HTMLInputElement;
    searchQuery = target.value;
    
    // Debounce search to avoid too many API calls
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
      dispatch('search', searchQuery);
    }, 300);
  }
  
  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      dispatch('search', searchQuery);
    }
  }
  
  function clearSearch() {
    searchQuery = '';
    dispatch('search', '');
  }
</script>

<div class="search-container">
  <div class="search-bar">
    <div class="search-icon">
      <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="11" cy="11" r="8"/>
        <path d="m21 21-4.35-4.35"/>
      </svg>
    </div>
    
    <input
      type="text"
      placeholder="Search applications..."
      bind:value={searchQuery}
      on:input={handleInput}
      on:keydown={handleKeydown}
      class="search-input"
      aria-label="Search applications"
    />
    
    {#if searchQuery}
      <button
        type="button"
        on:click={clearSearch}
        class="clear-button"
        aria-label="Clear search"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <line x1="18" y1="6" x2="6" y2="18"/>
          <line x1="6" y1="6" x2="18" y2="18"/>
        </svg>
      </button>
    {/if}
  </div>
</div>

<style>
  .search-container {
    margin-bottom: 24px;
  }
  
  .search-bar {
    position: relative;
    display: flex;
    align-items: center;
    background: #0d1117;
    border: 1px solid #30363d;
    border-radius: 6px;
    padding: 0 12px;
    transition: all 0.2s ease;
  }
  
  .search-bar:focus-within {
    border-color: #58a6ff;
    box-shadow: 0 0 0 3px rgba(56, 139, 253, 0.15);
  }
  
  .search-icon {
    color: #8b949e;
    margin-right: 8px;
    flex-shrink: 0;
  }
  
  .search-input {
    flex: 1;
    border: none;
    background: transparent;
    font-size: 14px;
    padding: 12px 0;
    color: #c9d1d9;
    outline: none;
  }
  
  .search-input::placeholder {
    color: #8b949e;
  }
  
  .clear-button {
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
  }
  
  .clear-button:hover {
    background: #21262d;
    color: #c9d1d9;
  }
  
  .clear-button:focus {
    outline: 2px solid #58a6ff;
    outline-offset: 2px;
  }
  
  @media (max-width: 768px) {
    .search-input {
      font-size: 14px;
      padding: 10px 0;
    }
  }
</style>
