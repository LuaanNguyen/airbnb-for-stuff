<script lang="ts">
  import {onMount} from 'svelte'

  interface User {
    id: number; 
    email: string;
    phone_number: string;
    first_name: string;
    last_name: string; 
    nick_name: string;
    password: string;
  }

  let users: User[] = [];
  let loading = true 
  let error: string | null = null; 

  onMount(async ()=> {
    try {
      const response = await fetch('http://localhost:8080/api/users')
      if (!response.ok) {
        throw new Error('Failed to fetch items')
      }
      const data = await response.json(); 
      
      users = data.slice(0, 50); //limit to 50 users per fetch
    } catch (e) {
      error = e instanceof Error ? e.message : 'An Error occured while fetching for users data'
    } finally {
      loading = false
    }
  })
</script>


<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <h1 class="text-3xl font-bold mb-8">AirBnb for Stuff 📦</h1>

  <h2 class="text-2xl font-bold ">All Users</h2>
  {#if loading}
    <div class="text-center py-12">
      <div class="text-gray-600">Loading items...</div>
    </div>
  {:else if error}
    <div class="text-center py-12">
      <div class="text-red-600">{error}</div>
    </div>
  {:else}
    <div class="flex flex-col gap-3">
      {#each users as user}
        {#if user}
          <div class="flex flex-row justify-between gap-4 border p-2 items-center overflow-hidden"> 
              <p>
                {user.first_name}{' '}{user.last_name}
              </p>
              <p>
                {#if user.nick_name}
                ({user.nick_name})
                {/if}
              </p>
              <p class=" text-green-600">
                {user.phone_number}
              </p>
              <p class="truncate">
                {user.email}
              </p>
 
          </div>
        {/if}
      {/each}
    </div>
  {/if}

  
    <p class="mt-20">
      Visit <a class="text-blue-600 underline" href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation.
    </p>
  </div>

  

  