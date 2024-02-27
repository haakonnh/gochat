<script lang="js">
    import { currentUser, pb } from "./pocketbase";

    let username = null;
    let password = null;

    async function login() {
        await pb.collection('users').authWithPassword(username, password);
    }

    async function signUp() {
        try {
            const data = {
                username,
                password,
                passwordConfirm: password,
                name: 'Hi mom'
            };
            const createdUser = await pb.collection('users').create(data);
            await login();
        } catch (e) {
            console.error(e);
        }
    }
    
    function logout() {
        pb.authStore.clear();
    }

</script>

{#if $currentUser}
<div class="isLoggedIn">
    <p>
        Signed in as {$currentUser.username}
    </p>
    <button class="logout-button"  on:click={logout}>Logout</button>
    </div>
{:else}
    <div class="login-box">
        <div class="background">
            <div class="shape"></div>
            <div class="shape"></div>
        </div>
        <form class="loginform" on:submit|preventDefault>   
            <h3><span style="color: #23a2f6;">go</span><span style="color: #f09819;">chat</span></h3>
            <label for="username">Username</label>
            <input placeholder="Username" type="text" bind:value={username} id="username"/><br>
            <label for="password">Password</label>
            <input placeholder="Password" type="password" bind:value={password} id="password"/><br>
            <button on:click={login}>Login</button><br>
            <button on:click={signUp}>Sign Up</button><br>
        </form>
    </div>
{/if}