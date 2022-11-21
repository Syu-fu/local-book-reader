import React, { useState, useEffect, SyntheticEvent } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Autocomplete from "@mui/material/Autocomplete";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { SubmitHandler, Controller, useForm } from "react-hook-form";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type Tag from '../types/Tag'
import type BookGroup from '../types/BookGroup'
import { IPADDRESS, MINIO_PORT } from '../config/index'

interface BookGroupEditInput {
  bookId: string
  title: string
  titleReading: string
  author: string
  authorReading: string
  tagId: string[]
  tags: Tag[]
}

const searchCharacters = async (search: string) => {
  return axios.get<Tag[]>(
    `/tag/search/q=${search}`
  )
    .then((response) => {
      return response.data
    }).catch(error => {
      console.error(error);
      return undefined;
    });
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

const BookGroupEditPage = () => {
  const [searchString, setSearchString] = useState('');
  const [searchBookGroupString, setSearchBookGroupString] = useState('');
  const { control, handleSubmit, register, setValue, reset } = useForm<BookGroupEditInput>({ mode: 'onChange', defaultValues: { tags: [] } })
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [resultOptions, setResultOptions] = useState<Tag[]>([]);
  const [options, setOptions] = useState<BookGroup[]>([]);
  const [bookId, setBookId] = useState('');
  const [title, setTitle] = useState('');
  const [titleReading, setTitleReading] = useState('');
  const [author, setAuthor] = useState('');
  const [authorReading, setAuthorReading] = useState('');
  const [tags, setTags] = useState<Tag[]>([]);
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
  const onSubmit: SubmitHandler<any> = (data: BookGroupEditInput) => {
    setIsLoading(true)
    axios.put<any>(
      `/bookgroup/${data.bookId}`, {
      title: data.title,
      titleReading: data.titleReading,
      author: data.author,
      authorReading: data.authorReading,
      tags: data.tags
    }
    ).then((response) => {
      if (response.status === 201) {
        setApiMessage(`BookGroup ${response.data.title} has been updated.`)
        setApiError(false)
        uploadFile(response.data.bookId);
      }
      else {
        setApiMessage('Update failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Update failed. ${error}`)
      setApiError(true)
    })
    setIsLoading(false)
    console.log({
      title: data.title,
      titleReading: data.titleReading,
      author: data.author,
      authorReading: data.authorReading,
      tags: data.tags
    }
    )
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
        title: selectedValue.title,
        titleReading: selectedValue.titleReading,
        author: selectedValue.author,
        authorReading: selectedValue.authorReading,
        tags: selectedValue.tags
      })
    }
  };


  const debouncedSearchTerm = useDebounce(searchString, 300);
  useEffect(
    () => {
      if (debouncedSearchTerm !== "") {
        searchCharacters(debouncedSearchTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setResultOptions([...searchResults]);
          }
        });
      }
    },
    [debouncedSearchTerm]
  );

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
              {...register('title')}
              onChange={(event) => { return setTitle(event.target.value) }}
            />
            <TextField
              id="standard-search"
              label="Title Reading"
              type="search"
              variant="standard"
              value={titleReading}
              {...register('titleReading')}
              onChange={(event) => { return setTitleReading(event.target.value) }}
            />
            <TextField
              id="standard-search"
              label="Author"
              type="search"
              variant="standard"
              value={author}
              {...register('author')}
              onChange={(event) => { return setAuthor(event.target.value) }}
            />
            <TextField
              id="standard-search"
              label="Author Reading"
              type="search"
              variant="standard"
              value={authorReading}
              {...register('authorReading')}
              onChange={(event) => { return setAuthorReading(event.target.value) }}
            />
            <Stack direction="row" spacing={2} justifyContent="center">
              <Typography>
                Thumbnail
              </Typography>
              <Button variant="contained" component="label">
                Upload
                <input hidden accept="image/*" type="file" id='thumbnail' />
              </Button>
            </Stack>
            <Controller
              control={control}
              name="tagId"
              render={() => {
                return (
                  <Autocomplete
                    multiple
                    options={resultOptions}
                    value={tags}
                    onChange={(event, item) => {
                      const ids = item.map((i) => { return i.tagId })
                      const tagList: Tag[] = [];
                      ids.map((id) => {
                        const tag: Tag = { tagId: id, tagName: "" }
                        return tagList.push(tag)
                      })
                      setValue('tags', tagList)
                      setTags(item)
                    }}
                    getOptionLabel={(option) => { return option.tagName }}
                    renderInput={(params) => {
                      return (
                        <TextField
                          {...params}
                          id="standard-search"
                          label="Tags"
                          type="search"
                          variant="standard"
                          onChange={(event) => { return setSearchString(event.target.value) }}
                        />
                      )
                    }}
                  />
                )
              }}
            />
            <LoadingButton
              size="large"
              loadingIndicator="Edit"
              variant="outlined"
              loading={isLoading}
              onClick={handleSubmit(onSubmit)}
            >
              Edit
            </LoadingButton>
            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default BookGroupEditPage
