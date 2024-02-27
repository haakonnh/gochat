<script>
    import { onMount, onDestroy } from "svelte";
    import { pb, currentUser } from "./pocketbase";

    let hasAccess = false;
    let chatrooms;

    async function chatroomAccess() {
        try {
            const res = await pb.collection('chatrooms').getList(1, 50, {
                filter: `users ~ "${$currentUser.id}"`,
                expand: 'users',
            });
            hasAccess = res.items.length > 0;
            chatrooms = res.items;

            console.log("Chatrooms:", chatrooms);

        }
        catch (e) {
            console.error("Errrs:",e);
        }
    }

    onMount(async () => {
        chatroomAccess();
    });
    
</script>
{#if $currentUser && hasAccess}
    <h1>Dashboard</h1>
    <div class="dashboard-flex">
        {#each chatrooms as chatroom (chatroom.id)}
            <div class="card">
                <h2>{chatroom.name}</h2>
                <a href="/chatroom/{chatroom.id}">Join</a>
            </div>
        {/each}
    </div>

{:else}
    <h1>Access Denied</h1>
{/if}




