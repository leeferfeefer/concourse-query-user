<template>
  <v-container>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <transition name="slide" mode="out-in">
          <v-card v-if="!paid" key="toll" class="elevation-12">
            <v-toolbar dark :color="error ? 'red' : 'blue'">
              <v-toolbar-title>{{error ? 'Uh-oh!' : 'Enter Credentials'}}</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form
                      ref="form"
                      v-model="valid"
                      lazy-validation
              >
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
                <v-btn
                      :disabled="!valid"
                      @click="payTheTrollToll"
                >
                  Pay the troll toll
                </v-btn>
                <label v-if="error" color="red">Network Error - Please try again</label>
              </v-form>
            </v-card-text>
          </v-card>
          <v-card v-if="paid" key="success" class="elevation-12">
            <v-toolbar dark color="green">
              <v-toolbar-title>Success</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              Toll Paid
            </v-card-text>
          </v-card>
        </transition>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'QueryUser',
    props: {},
    data: () => ({
      valid: true,
      paid: false,
      error: false,
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
      async payTheTrollToll() {
        if (this.$refs.form.validate()) {
          const payload = {
            data: {
              username: btoa(this.username),
              password: btoa(this.password)
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
            await instance.post('/toll', payload);
            this.error = false;
            this.paid = true;
          } catch (err) {
            this.error = true;
          }
        }
      }
    }
  }
</script>

<style scoped>

  .slide-enter-active, .slide-leave-active {
    transition: opacity .2s ease-in-out, transform .2s ease-in-out;
  }

  .slide-enter, .slide-leave-to {
    opacity: 0;
    transform: translateX(-20px);
  }
</style>
