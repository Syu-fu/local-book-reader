import React from "react";
import { render, screen } from "@testing-library/react";
import BookListItem from "./components/BookListItem";
import type Book from "./types/Book";

test("renders booklist item", () => {
  const book: Book = {
    bookId: "",
    volume: "1",
    displayOrder: "1",
    thumbnail: "",
    title: "The Hitchhiker's Guide to the Galaxy",
    filepath: "Douglas Adams",
    author: "Douglas Adams",
    publisher: "",
  };
  render(<BookListItem src={book.thumbnail} title={book.title} author={book.author} volume={book.volume} />);
  const titleElement = screen.getByText(
    /The Hitchhiker's Guide to the Galaxy/i
  );
  expect(titleElement).toBeInTheDocument();
  const authorElement = screen.getByText(/Douglas Adams/i);
  expect(authorElement).toBeInTheDocument();
});
