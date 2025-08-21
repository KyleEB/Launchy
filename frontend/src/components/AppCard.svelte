<script lang="ts">
  import type { AppInfo } from '../types/app';
  
  const { app, isSelected, onLaunch } = $props<{
    app: AppInfo;
    isSelected: boolean;
    onLaunch: (exec: string) => void;
  }>();
  
  function handleLaunch() {
    onLaunch(app.exec);
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      handleLaunch();
    }
  }
  

</script>

<div 
  on:click={handleLaunch} 
  on:keydown={handleKeydown} 
  role="button" 
  tabindex="0"
  class="border transition-all duration-200 cursor-pointer focus:outline-none focus:ring-2 focus:ring-opacity-50 {isSelected ? 'ring-2 ring-yellow-400 ring-opacity-75' : ''}"
  style="border-radius: 16px; padding: 12px; background-color: var(--color-card-background); border-color: var(--color-border);"
>
    <div class="mb-1">
    <h3 class="text-sm font-medium truncate">{app.name}</h3>
  </div>
  
  {#if app.comment}
    <p class="text-xs mb-1 line-clamp-1" style="color: var(--color-secondary-text);">{app.comment}</p>
  {/if}
  
  {#if app.categories}
    <div class="flex flex-wrap gap-1 mb-1">
      {#each app.categories.split(';').slice(0, 2) as category}
        {#if category.trim()}
          <span class="inline-block px-1.5 py-0.5 text-xs font-medium rounded border" style="background-color: #21262d; color: var(--color-primary-text); border-color: var(--color-border);">{category.trim()}</span>
        {/if}
      {/each}
    </div>
  {/if}
  
  <p class="text-xs text-[#a855f7] font-mono truncate">{app.exec}</p>
</div>


