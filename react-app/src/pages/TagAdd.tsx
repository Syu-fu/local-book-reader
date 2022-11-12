import React, { useState, useEffect } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import { SubmitHandler, useForm } from "react-hook-form";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type Tag from '../types/Tag'

interface TagAddInput {
  tagName: string
}

const searchCharacters = async (search: string) => {
  return axios.get<Tag[]>(
    `/tag/search/name/${search}`
  )
    .then((response) => {
      return response.data
    })
    .catch(error => {
      console.error(error);
      return undefined;
    });

}

const TagAddPage = () => {
  const [tagName, setTagName] = useState('');
  const [loading, setLoading] = useState(true);
  const [results, setResults] = useState<Tag[]>([]);
  const { register, handleSubmit } = useForm<TagAddInput>()
  const [helperText, setHelperText] = useState('Please enter at least 1 character')
  const [isErrorTagName, setIsErrorTagName] = useState(true) 
  const [apiMessage, setApiMessage] = useState('') 
  const [apiError, setApiError] = useState(false) 
  const onSubmit: SubmitHandler<TagAddInput> = (data) => {
    setLoading(true)
    axios.post<any>(
      '/tag/', {
      tagName: data.tagName
    }
    ).then((response) => {
      if (response.status === 201){
        setApiMessage(`tagname ${response.data.tagName} has been registered.`)
        setApiError(false)
        setIsErrorTagName(true)
        setTagName('')
      }
      else{
        setApiMessage('Registration failed.')
        setApiError(true)
      }
    }).catch((error) => {
        setApiMessage(`Registration failed. ${error}`)
        setApiError(true)
    })
  }

  const debouncedSearchTerm = useDebounce(tagName, 300);
  useEffect(
    () => {
      if (debouncedSearchTerm !== "") {
        searchCharacters(debouncedSearchTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setResults(searchResults);
            if (results.length === 0) {
              setLoading(false);
              setIsErrorTagName(false)
              setHelperText('')
            }
            else {
              setLoading(true)
              setIsErrorTagName(true)
              setHelperText('The same tag name is registered')
            }
          }
        });
      }
      else {
        setLoading(true)
        setHelperText('Please enter at least 1 character')
        setIsErrorTagName(true)
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
              label="Tag name"
              type="search"
              variant="standard"
              value={tagName}
              {...register('tagName')}
              onChange={(event) => { return setTagName(event.target.value) }}
              error={isErrorTagName}
              helperText={helperText}
            />
            <LoadingButton
              size="large"
              loading={loading}
              loadingIndicator="Add"
              variant="outlined"
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

export default TagAddPage
