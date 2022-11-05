import { AxiosError } from 'axios'
import { useEffect, useState } from 'react'
import axios from '../lib/axios'
import type BookResponse from '../types/BookResponse'
import type Book from '../types/Book'


export const useFetchBook = (bookIdParams: string | undefined, volumeParams: string | undefined) => {
  const [res, setRes] = useState<BookResponse>({ data: null, error: null, loading: false })

  const fetchRequest = (bookId: string | undefined, volume: string | undefined) => {
    setRes(prevState => { return { ...prevState, loading: true } })
    axios.get<Book>(`/book/${bookId}/${volume}`, {
    }).then((response) => {
      setRes({ data: response.data, error: null, loading: false })
    }).catch((error: AxiosError) => {
      console.log(error);
      setRes({ data: null, error, loading: false })
    })
  }

  useEffect(() => {
    fetchRequest(bookIdParams, volumeParams)
  }, [])
  return res
}

export default useFetchBook
