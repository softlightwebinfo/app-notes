import { writable } from "svelte/store";
const notes = writable([]);
const noteSelection = writable(null);

export const useNotes = () => {
  return { notes, noteSelection };
};
