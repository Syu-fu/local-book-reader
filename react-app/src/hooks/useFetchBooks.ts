import { AxiosError } from 'axios'
import { useEffect, useState } from 'react'
import axios from '../lib/axios'
import type BooksResponse from '../types/BooksResponse'
import type Book from '../types/Book'


export const useFetchBooks = (bookIdParams: string | undefined) => {
  const [res, setRes] = useState<BooksResponse>({ data: null, error: null, loading: false })

  const fetchRequest = (bookId: string | undefined) => {
    setRes(prevState => {return { ...prevState, loading: true }})
    axios.get<Book[]>(`/book/${bookId}`, {
    }).then((response) => {
      setRes({ data: response.data, error: null, loading: false })
    }).catch((error: AxiosError) => {
      console.log(error);
      setRes({ data: null, error, loading: false })
    })
  }

  useEffect(() => {
    fetchRequest(bookIdParams)
  }, [])
  return res
}

export default useFetchBooks
