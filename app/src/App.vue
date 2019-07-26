<template>
    <v-app>
        <v-toolbar app>
            <v-toolbar-title class="headline">
                <span>Bridge Troll</span>
            </v-toolbar-title>
        </v-toolbar>
        <v-content>
            <QueryUser/>
        </v-content>
    </v-app>
</template>

<script>
    import QueryUser from './components/QueryUser'

    export default {
        name: 'App',
        components: {
            QueryUser
        },
        data: () => ({})
    }
</script>
