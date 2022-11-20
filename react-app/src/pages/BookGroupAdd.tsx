import React, { useState, useEffect } from 'react'
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
import { IPADDRESS, MINIO_PORT } from '../config/index'

interface BookGroupAddInput {
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

const BookGroupAddPage = () => {
  const [searchString, setSearchString] = useState('');
  const { control, handleSubmit, register, setValue } = useForm<BookGroupAddInput>({ mode: 'onChange', defaultValues: { tags: [] } })
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [resultOptions, setResultOptions] = useState<Tag[]>([]);
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
  const onSubmit: SubmitHandler<any> = (data: BookGroupAddInput) => {
    setIsLoading(true)
    axios.post<any>(
      `/bookgroup/`, {
      title: data.title,
      titleReading: data.titleReading,
      author: data.author,
      authorReading: data.authorReading,
      tags: data.tags
    }
    ).then((response) => {
      if (response.status === 201) {
        setApiMessage(`BookGroup ${response.data.title} has been created.`)
        setApiError(false)
        uploadFile(response.data.bookId);
      }
      else {
        setApiMessage('Create failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Create failed. ${error}`)
      setApiError(true)
    })
    setIsLoading(false)
  }


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


  return (
    <div>
      <ResponsiveDrawer>
        <Container maxWidth="sm" sx={{ pt: 5 }}>
          <Stack spacing={3}>
            <TextField
              id="standard-search"
              label="Title"
              type="search"
              variant="standard"
              {...register('title')}
            />
            <TextField
              id="standard-search"
              label="Title Reading"
              type="search"
              variant="standard"
              {...register('titleReading')}
            />
            <TextField
              id="standard-search"
              label="Author"
              type="search"
              variant="standard"
              {...register('author')}
            />
            <TextField
              id="standard-search"
              label="Author Reading"
              type="search"
              variant="standard"
              {...register('authorReading')}
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
                    onChange={(event, item) => {
                      const ids = item.map((i) => { return i.tagId })
                      const tags: Tag[] = [];
                      ids.map((id) => {
                        const tag: Tag = { tagId: id, tagName: "" }
                        return tags.push(tag)
                      })
                      setValue('tags', tags)
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
              Add
            </LoadingButton>
            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default BookGroupAddPage
