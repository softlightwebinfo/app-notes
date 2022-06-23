<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import { Greet } from "../wailsjs/go/main/App.js";
  import { Router, Route, Link } from "svelte-navigator";
  import NotesHome from "./routers/NotesHome.svelte";

  let resultText: string = "Please enter your name below 👇";
  let name: string;

  function greet(): void {
    Greet(name).then((result) => (resultText = result));
  }
</script>

<main>
  <Router>
    <Route path="/">
      <NotesHome />
    </Route>
  </Router>
</main>

<style lang="scss">
</style>
