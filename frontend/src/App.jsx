import React, { useState, useEffect } from 'react';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';

function App() {
  const [url, setUrl] = useState('');
  const [nLink, setNLink] = useState('');

  useEffect(() => {
    const slug = window.location.pathname.slice(1);
    if (!slug) return;

    fetch(`http://localhost:8080/${slug}`)
      .then(res => {
        if (!res.ok) throw new Error('Slug not found');
        return res.json();
      })
      .then(data => {
        if (data.url) {
          // front-end redirect
	  let dest = data.url;
    	  // if it doesn’t already have http:// or https://, add https://
    	  if (!/^https?:\/\//i.test(dest)) {
      		dest = 'https://' + dest;
    	  }
    	  window.location.href = dest;
        }
      })
      .catch(err => {
        console.error('Redirect error:', err);
        // you could set some state to show "Not found" here
      });
  }, []);

  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url }), // or whatever your server expects
      });

      if (!response.ok) {
	      console.log(await response.json());
      }
      const data = await response.json();
      console.log(data);
      setNLink('http://localhost:5173/'+data.hashslug); // update based on actual server response field
    } catch (err) {
      console.error('Error:', err);
      setNLink('');
    }
  };

  return (
    <Paper
      sx={{
        width: '100vw',
        height: '100vh',
        p: 4,
        boxSizing: 'border-box',
        display: 'flex',
        flexDirection: 'column',
      }}
    >

	<Typography variant='h1' sx={{ textAlign: 'left', mb: 4}}>
		URL King
	</Typography>
        <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>
          Enter a URL to get started
        </Typography>

      <Box
        sx={{
          flex: 1,
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          gap: 2,
        }}
      >
        <TextField
          id="link"
          label="URL"
          variant="outlined"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          sx={{ width: 400 }}
        />
        <Button variant="contained" onClick={handleSubmit}>
          Shorten
        </Button>

        {nLink && (
          <Typography variant="body1" sx={{ mt: 2 }}>
            Shortened Link:
            <a href={nLink} target="_blank" rel="noopener noreferrer">
              {nLink}
            </a>
          </Typography>
        )}
      </Box>

      <Typography variant="body2" sx={{ textAlign: 'center', mt: 'auto' }}>
        © 2025 Kieran Fane
      </Typography>
    </Paper>
  );
}

export default App;

