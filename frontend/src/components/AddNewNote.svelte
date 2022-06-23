<form>
  <h1>No hay notas</h1>
  <input type="text" name="title" placeholder="Escribe el titulo de la nota" />
  <button>Crear nueva nota</button>
</form>

<style lang="scss">
  form {
    display: flex;
    height: 100%;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  h1 {
    margin-bottom: 40px;
    font-size: 3em;
  }
  input {
    width: 100%;
    height: 50px;
    border-radius: 80px;
    border: 0;
    outline: 0;
    padding: 10px;
    max-width: 400px;
  }
  button {
    padding: 10px;
    border-radius: 20px;
    border: 0;
    background-color: white;
    margin-top: 40px;
    font-size: 20px;
    padding-left: 40px;
    padding-right: 40px;
  }
</style>
