import { AxiosError } from 'axios'
import Book from './Book'

type BooksResponse = {
  data: Book[] | null;
  error: AxiosError | null;
  loading: boolean
}

export default BooksResponse
