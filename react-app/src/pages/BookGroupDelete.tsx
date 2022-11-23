import React, { useState, useEffect, SyntheticEvent } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Autocomplete from "@mui/material/Autocomplete";
import Typography from "@mui/material/Typography";
import { SubmitHandler, useForm } from "react-hook-form";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type Tag from '../types/Tag'
import type BookGroup from '../types/BookGroup'
import { IPADDRESS, MINIO_PORT } from '../config/index'

interface BookGroupDeleteInput {
  bookId: string
}

const searchBookGroup = async (search: string) => {
  return axios.get<BookGroup[]>(
    `/bookgroup/search/q=${search}`
  )
    .then((response) => {
      return response.data
    }).catch(error => {
      console.error(error);
      return undefined;
    });
}

const BookGroupDeletePage = () => {
  const [searchBookGroupString, setSearchBookGroupString] = useState('');
  const { handleSubmit, register, reset } = useForm<BookGroupDeleteInput>({ mode: 'onChange' })
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [options, setOptions] = useState<BookGroup[]>([]);
  const [bookId, setBookId] = useState('');
  const [title, setTitle] = useState('');
  const [titleReading, setTitleReading] = useState('');
  const [author, setAuthor] = useState('');
  const [authorReading, setAuthorReading] = useState('');
  const [tags, setTags] = useState<Tag[]>([]);
  const resultOptions : Tag[] = []
  const uploadFile = (id: string) => {
    const input = document.getElementById('thumbnail') as HTMLInputElement | null;

    if (input != null) {
      const file = input.files ? input.files[0] : null;
      if (!file) return;

      axios.put<any>(
        `http://${IPADDRESS}:${MINIO_PORT}/data/thumbnail/bookgroup/${id}`,
        file
      )
    }
  }
  const onSubmit: SubmitHandler<any> = (data: BookGroupDeleteInput) => {
    setIsLoading(true)
    axios.delete<any>(
      `/bookgroup/${data.bookId}`, {
    }
    ).then((response) => {
      if (response.status === 200) {
        setApiMessage(`BookGroup ${response.data.title} has been deleted.`)
        setApiError(false)
        uploadFile(response.data.bookId);
      }
      else {
        setApiMessage('Delete failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Delete failed. ${error}`)
      setApiError(true)
    })
    setIsLoading(false)
  }

  const selectedChange = (event: SyntheticEvent, selectedValue: BookGroup | null) => {// eslint-disable-line no-unused-vars
    if (selectedValue !== null) {
      setBookId(selectedValue.bookId)
      setTitle(selectedValue.title)
      setTitleReading(selectedValue.titleReading)
      setAuthor(selectedValue.author)
      setAuthorReading(selectedValue.authorReading)
      setTags([...tags, ...selectedValue.tags])
      reset({
        bookId: selectedValue.bookId,
      })
    }
  };


  const debouncedBookGroupSearchTerm = useDebounce(searchBookGroupString, 300);
  useEffect(
    () => {
      if (debouncedBookGroupSearchTerm !== "") {
        searchBookGroup(debouncedBookGroupSearchTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setOptions([...searchResults]);
          }
        });
      }
    },
    [debouncedBookGroupSearchTerm]
  );

  return (
    <div>
      <ResponsiveDrawer>
        <Container maxWidth="sm" sx={{ pt: 5 }}>
          <Stack spacing={3}>
            <Autocomplete
              options={options}
              onChange={(event, value) => { return selectedChange(event, value) }}
              getOptionLabel={(option) => { return option.title }}
              renderInput={(params) => {
                return (
                  <TextField
                    {...params}
                    id="standard-search"
                    label="Search"
                    type="search"
                    variant="standard"
                    onChange={(event) => { return setSearchBookGroupString(event.target.value) }}
                  />
                )
              }}
            />
            <TextField
              id="standard-search"
              label="ID"
              type="search"
              variant="standard"
              value={bookId}
              disabled
              {...register('bookId')}
            />
            <TextField
              id="standard-search"
              label="Title"
              type="search"
              variant="standard"
              value={title}
              disabled
            />
            <TextField
              id="standard-search"
              label="Title Reading"
              type="search"
              variant="standard"
              value={titleReading}
              disabled
            />
            <TextField
              id="standard-search"
              label="Author"
              type="search"
              variant="standard"
              value={author}
              disabled
            />
            <TextField
              id="standard-search"
              label="Author Reading"
              type="search"
              variant="standard"
              value={authorReading}
              disabled
            />
            <Autocomplete
              multiple
              disabled
              options={resultOptions}
              value={tags}
              getOptionLabel={(option) => { return option.tagName }}
              renderInput={(params) => {
                return (
                  <TextField
                    {...params}
                    id="standard-search"
                    label="Tags"
                    type="search"
                    variant="standard"
                  />
                )
              }}
            />
            <LoadingButton
              size="large"
              loadingIndicator="Delete"
              variant="outlined"
              loading={isLoading}
              onClick={handleSubmit(onSubmit)}
            >
              Delete
            </LoadingButton>
            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default BookGroupDeletePage
