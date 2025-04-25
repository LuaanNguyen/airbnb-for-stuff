<script lang="ts">
  import { onMount } from 'svelte';
  import { getAvailableItems, getAllCategories, searchItems } from '$lib/services/api';
  import { isAuthenticated } from '$lib/auth';
  import { goto } from '$app/navigation';
  import type { ItemWithOwner, Category, SearchParams } from '$lib/types';

  let items: ItemWithOwner[] = [];
  let categories: Category[] = [];
  let loading = true;
  let error: string | null = null;

  // Search params
  let searchQuery = '';
  let selectedCategory: number | null = null;
  let minPrice: number | null = null;
  let maxPrice: number | null = null;
  let showOnlyAvailable = true;

  onMount(async () => {
    // Check if user is authenticated
    if (!$isAuthenticated) {
      goto('/login');
      return;
    }

    try {
      // Load data
      await Promise.all([
        loadItems(),
        loadCategories()
      ]);
    } catch (e) {
      error = e instanceof Error ? e.message : 'An error occurred';
    } finally {
      loading = false;
    }
  });

  async function loadItems() {
    try {
      if (searchQuery || selectedCategory || minPrice || maxPrice) {
        // Use search API if any filter is applied
        const searchParams: SearchParams = {
          query: searchQuery || undefined,
          categoryID: selectedCategory || undefined,
          minPrice: minPrice || undefined,
          maxPrice: maxPrice || undefined,
          available: showOnlyAvailable
        };
        
        const searchResults = await searchItems(searchParams);
        items = searchResults as unknown as ItemWithOwner[]; // Type cast as API returns similar structure
      } else {
        // Otherwise load all available items
        items = await getAvailableItems();
      }
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to load items';
    }
  }

  async function loadCategories() {
    try {
      categories = await getAllCategories();
    } catch (e) {
      console.error('Failed to load categories:', e);
    }
  }

  async function handleSearch() {
    loading = true;
    error = null;
    await loadItems();
    loading = false;
  }

  function resetFilters() {
    searchQuery = '';
    selectedCategory = null;
    minPrice = null;
    maxPrice = null;
    showOnlyAvailable = true;
    handleSearch();
  }

  // Format price to USD
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2
    }).format(price / 100); // Converting cents to dollars
  };
</script>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <h1 class="text-3xl font-bold text-gray-900 mb-8">Available Items</h1>
  
  <!-- Search and filters -->
  <div class="bg-white p-4 rounded-lg shadow-md mb-6">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
      <!-- Search query -->
      <div>
        <label for="search" class="block text-sm font-medium text-gray-700 mb-1">Search</label>
        <input
          type="text"
          id="search"
          bind:value={searchQuery}
          placeholder="Search items..."
          class="w-full px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>
      
      <!-- Category filter -->
      <div>
        <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
        <select
          id="category"
          bind:value={selectedCategory}
          class="w-full px-3 py-2 border border-gray-300 rounded-md"
        >
          <option value={null}>All Categories</option>
          {#each categories as category}
            <option value={category.id}>{category.name}</option>
          {/each}
        </select>
      </div>
      
      <!-- Price range -->
      <div>
        <label for="minPrice" class="block text-sm font-medium text-gray-700 mb-1">Min Price ($)</label>
        <input
          type="number"
          id="minPrice"
          bind:value={minPrice}
          min="0"
          placeholder="Min price"
          class="w-full px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>
      
      <div>
        <label for="maxPrice" class="block text-sm font-medium text-gray-700 mb-1">Max Price ($)</label>
        <input
          type="number"
          id="maxPrice"
          bind:value={maxPrice}
          min="0"
          placeholder="Max price"
          class="w-full px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>
    </div>
    
    <div class="flex items-center justify-between">
      <label class="flex items-center">
        <input type="checkbox" bind:checked={showOnlyAvailable} class="mr-2" />
        <span class="text-sm">Show only available items</span>
      </label>
      
      <div class="flex space-x-2">
        <button
          on:click={resetFilters}
          class="px-4 py-2 text-sm border border-gray-300 rounded-md hover:bg-gray-50"
        >
          Reset
        </button>
        <button
          on:click={handleSearch}
          class="px-4 py-2 text-sm bg-blue-600 text-white rounded-md hover:bg-blue-700"
        >
          Search
        </button>
      </div>
    </div>
  </div>
  
  {#if loading}
    <div class="text-center py-12">
      <div class="text-gray-600">Loading items...</div>
    </div>
  {:else if error}
    <div class="text-center py-12">
      <div class="text-red-600">{error}</div>
    </div>
  {:else if items.length === 0}
    <div class="text-center py-12">
      <div class="text-gray-600">No items found. Try different search criteria.</div>
    </div>
  {:else}
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      {#each items as item}
        <div class="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-200 hover:-translate-y-1 hover:shadow-lg">
          <div class="w-full h-48 flex items-center justify-center bg-gray-100 text-gray-500">
            No image available
          </div>
          
          <div class="p-4">
            <h2 class="text-xl font-semibold text-gray-900 mb-2 truncate">
              {item.name}
            </h2>
            
            <p class="text-2xl font-bold text-green-600 mb-2">
              {formatPrice(item.price)}
            </p>
            
            <p class="text-gray-600 text-sm mb-3 line-clamp-2">
              {item.description}
            </p>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-500">
                Owner: {item.owner_name}
              </span>
              <button 
                on:click={() => goto(`/items/rent/${item.id}`)}
                class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm font-medium transition-colors duration-200"
              >
                Rent Now
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div> 