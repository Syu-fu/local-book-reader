import { AxiosError } from 'axios'
import Book from './Book'

type BookResponse = {
  data: Book | null;
  error: AxiosError | null;
  loading: boolean
}

export default BookResponse
