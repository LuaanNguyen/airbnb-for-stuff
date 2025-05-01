<script lang="ts">
  import { goto } from '$app/navigation';
  import { login } from '$lib/services/api';
  import { setAuth } from '$lib/auth';
  import type { LoginResponse } from '$lib/types';
  
  let email = 'alice@example.com';
  let password = 'anything';
  let error = '';
  let loading = false;

  async function handleLogin() {
    loading = true;
    error = '';
    
    try {
      const data = await login(email, password) as LoginResponse;
      
      // Set auth data using our helper function
      setAuth(data);
      
      // Redirect to items page
      goto('/items');
    } catch (err) {
      error = err instanceof Error ? err.message : 'An unknown error occurred';
    } finally {
      loading = false;
    }
  }
</script>

<div class="max-w-md mx-auto p-8">

  
  {#if error}
    <div class="bg-red-50 text-red-800 p-3 rounded-md mb-4">
      {error}
    </div>
  {/if}
  
  <form on:submit|preventDefault={handleLogin} class="pt-40">
    <h1 class="text-center text-2xl font-bold mb-6">Login to <br/> AirBnb for Items 📦</h1>
    <div class="mb-4">
      <label for="email" class="block mb-2 font-medium">Email</label>
      <input 
        type="email" 
        id="email" 
        bind:value={email} 
        required
        placeholder="Enter your email"
        class="w-full px-3 py-2 border border-gray-300 text-base focus:outline-none focus:ring-2 focus:ring-indigo-500"
      />
    </div>
    
    <div class="mb-4">
      <label for="password" class="block mb-2 font-medium">Password</label>
      <input 
        type="password" 
        id="password" 
        bind:value={password} 
        required
        placeholder="Enter your password"
        class="w-full px-3 py-2 border border-gray-300 text-base focus:outline-none focus:ring-2 focus:ring-indigo-500"
      />
    </div>

    
    <button 
      type="submit" 
      disabled={loading}
      class="w-full py-3 px-4 hover:bg-black hover:text-white border disabled:cursor-not-allowed"
    >
      {loading ? 'Logging in...' : 'Login'}
    </button>
  </form>
</div>
