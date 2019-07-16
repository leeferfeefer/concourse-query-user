<template>
  <v-form
    ref="form"
    v-model="valid"
    lazy-validation
  >
    <v-container>
      <v-layout align-start justify-start column fill-height>
        <v-flex xs12 sm6 md3>
          <v-text-field
                label="Username"
                outline
                v-model="username"
                required
                :rules="usernameRules"
          ></v-text-field>
          <v-text-field
                  type="password"
                  label="Password"
                  outline
                  v-model="password"
                  required
                  :rules="passwordRules"
          ></v-text-field>
        </v-flex>
        <v-flex xs12 sm6 md3>
          <div>
            <v-btn
              :disabled="!valid"
              @click="feed"
            >
              Feed
            </v-btn>
          </div>
        </v-flex>
      </v-layout>
    </v-container>
  </v-form>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'QueryUser',
    props: {},
    data: () => ({
      valid: true,
      username: '',
      usernameRules: [
        v => !!v || 'Username is required',
      ],
      password: '',
      passwordRules: [
        v => !!v || 'Password is required',
      ]
    }),
    methods: {
      feed() {
        if (this.$refs.form.validate()) {
          const payload = {
            data: {
              username: this.username,
              password: this.password
            }
          };

          const instance = axios.create({
            baseURL: 'https://localhost:3001/',
            timeout: 1000,
            headers: {
              "Content-Type": "application/json"
            }
          });

          try {
            instance.post('/substance', payload);
          } catch (err) {
            console.log("err")
          }
        }
      }
    }
  }
</script>

<style scoped>

</style>
