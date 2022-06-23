<script lang="ts">
  import AddNewNote from "./../components/AddNewNote.svelte";
  import { useNotes } from "../stores/notes";
  import Note from "../components/Note.svelte";
  const { notes, noteSelection } = useNotes();
</script>

<div class="section">
  <aside>
    {#if !$notes.length}
      <p>No hay notas disponibles</p>
    {/if}
    {#each $notes as note (note.id)}
      <Note {...note} />
    {/each}
  </aside>
  <section>
    {#if !$notes.length || !$noteSelection}
      <AddNewNote />
    {/if}
  </section>
</div>

<style lang="scss">
  h1 {
    margin: 0;
  }
  .section {
    display: flex;
    height: 100vh;
    gap: 20px;
  }
  aside {
    width: 300px;
    background: white;
    height: 100%;
    overflow: auto;
  }
  section {
    display: flex;
    flex-direction: column;
    flex: 1;
  }
  p {
    color: black;
    font-size: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>
