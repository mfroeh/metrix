import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Typography, Container, Box, Button, TextField, Select, MenuItem, FormControl, InputLabel } from '@mui/material';

const Home: React.FC = () => {
  const [summonerInput, setSummonerInput] = useState('');
  const [region, setRegion] = useState('EUW1');
  const navigate = useNavigate();

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const [name, tag] = summonerInput.split('#');
    if (name && tag) {
      navigate(`/summoner/${name}/${tag}`);
    } else {
      // Handle error: Invalid input format
      console.error('Invalid summoner name format');
    }
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 4 }}>
        <Typography variant="h2" component="h1" gutterBottom>
          Welcome to LoL Metrix
        </Typography>
        <Typography variant="body1" paragraph>
          Enter a summoner name and tag to get started.
        </Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            label="Summoner Name#Tag"
            variant="outlined"
            value={summonerInput}
            onChange={(e) => setSummonerInput(e.target.value)}
            placeholder="e.g. Faker#KR1"
            margin="normal"
          />
          <FormControl fullWidth margin="normal">
            <InputLabel id="region-select-label">Region</InputLabel>
            <Select
              labelId="region-select-label"
              value={region}
              label="Region"
              onChange={(e) => setRegion(e.target.value)}
            >
              <MenuItem value="EUW1">EUW</MenuItem>
              <MenuItem value="NA1">NA</MenuItem>
            </Select>
          </FormControl>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            fullWidth
            sx={{ mt: 2 }}
          >
            Search Summoner
          </Button>
        </form>
      </Box>
    </Container>
  );
};

export default Home;