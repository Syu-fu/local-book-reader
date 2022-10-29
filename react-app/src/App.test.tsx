import React from "react";
import { render, screen } from "@testing-library/react";
import BookGroupListItem from "./components/BookGroupListItem";
import type BookGroup from "./types/BookGroup";

test("renders bookgroup list item", () => {
  const bookgroup: BookGroup = {
    bookId: "",
    title: "The Hitchhiker's Guide to the Galaxy",
    titleReading: "",
    author: "Douglas Adams",
    authorReading: "",
    thumbnail: "",
    tags: [
      {
        TagId: "",
        TagName: "novel",
      },
    ],
  };
  render(<BookGroupListItem bookgroup={bookgroup} />);
  const titleElement = screen.getByText(
    /The Hitchhiker's Guide to the Galaxy/i
  );
  expect(titleElement).toBeInTheDocument();
  const authorElement = screen.getByText(/Douglas Adams/i);
  expect(authorElement).toBeInTheDocument();
  const tagElement = screen.getByText(/novel/i);
  expect(tagElement).toBeInTheDocument();
});
