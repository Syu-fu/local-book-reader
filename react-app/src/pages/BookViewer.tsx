import React, { useState, useEffect } from 'react'
import ComicViewer from 'react-comic-viewer'
import { useParams } from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Divider from '@mui/material/Divider'
import mc from '../lib/mc'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import { useFetchBook } from '../hooks/useFetchBook'

const BookViewer: React.FC = () => {
  const params = useParams<{ bookId: string, volume: string }>()
  const { data, error, loading } = useFetchBook(params.bookId, params.volume)
  const [filenames, setFilenames] = useState<string[]>([]);
  const bucketName = "data"
  const direction = data ? data.direction : "ltr"

  useEffect(
    () => {
      let ignore = false;
      const listObjectsOfBucket = async () => {
        try {
          const objectsStream = mc.extensions.listObjectsV2WithMetadata(bucketName, `file/${params.bookId}/${params.volume}/`, true, '')
          console.log(objectsStream)
          objectsStream.on('data', async (chunk) => {
            const { name: objectName } = chunk
            const presignedUrl = await mc.presignedGetObject(bucketName, objectName)
            if (!ignore) {
              setFilenames((pre) => {
                return [...pre, presignedUrl]
              })
            }
          });
        } catch (err) {
          console.log("Error in list objects", err)
        }
      };
      listObjectsOfBucket();
      return () => {
        ignore = true;
      }
    }, [setFilenames]
  );

  return (
    <div>
      <ResponsiveDrawer>
        <ComicViewer
          initialCurrentPage={0}
          initialIsExpansion={false}
          pages={
            filenames
          }
          direction={direction}
          switchingRatio={0.75}
          text={{
            expansion: "expantion",
            fullScreen: "fullscreen",
            move: "move",
            normal: "normal",
          }}
        />
        {error && <div>error</div>}
        {loading && <div>loading</div>}
        {data &&
          <div style={{ padding: "16px" }}>
            <Typography variant="h5">{data.title}</Typography>
            <Divider />
            <Typography variant="body1">{data.author}</Typography>
          </div>
        }
      </ResponsiveDrawer>
    </div>
  )

}
export default BookViewer
