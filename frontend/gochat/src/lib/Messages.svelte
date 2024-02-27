<script>
    import { onMount, onDestroy } from "svelte";
    import { pb, currentUser } from "./pocketbase";

    let newMessage;
    let messages = [];
    let unsubscribe;

    onMount(async () => {
        const resultList = await pb.collection('messages').getList(1, 50, {
            sort: 'created',
            expand: 'user',
        });
        messages = resultList.items;

        unsubscribe = await pb.collection('messages').subscribe('*', async ({action, record}) => {
            if (action === 'create') {
                const user = await pb.collection('users').getOne(record.user);
                record.expand = {user};
                messages = [...messages, record];
            }
            if (action === 'delete') {
                messages = messages.filter(message => message.id !== record.id);
            }
        });
    }); 

    onDestroy(() => {
        unsubscribe?.();
    });

    async function sendMessage() {
        const data = {
            text: newMessage,
            user: $currentUser.id,
        };
        const createdMessage = await pb.collection('messages').create(data);
        newMessage = '';
    }

    let hasAccess = false;

    
</script>




{#if $currentUser}
    <h1>Messages</h1>
    <div class="messages">
        {#each messages as message (message.id)}
            <div class="message">
                <img 
                class="avatar"
                src={`https://api.dicebear.com/7.x/pixel-art/svg?seed=${message.expand?.user?.username}`}
                alt="avatar"
                width="40px"
                />
            </div>
            <div>
                <small>
                    Sent by @{message.expand?.user?.username}
                </small>
                <p class="message-text">{message.text}</p>
            </div>
        {/each}
    </div>

    <form on:submit|preventDefault={sendMessage}>
        <input placeholder="Message" type="text" bind:value={newMessage}/>
        <button type="submit">Send</button>
    </form>
{:else}
    <!-- <p>
        Log in to join a chat room
    </p> -->
{/if}
