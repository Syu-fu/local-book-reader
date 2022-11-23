import type Tag from './Tag'

type BookGroup= {
  bookId:string;
  title: string;
  titleReading: string;
  author: string;
  authorReading: string;
  tags: Tag[];
}

export default BookGroup
