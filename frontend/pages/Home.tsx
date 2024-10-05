import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Typography, Container, Box, Button } from '@mui/material';

const Home: React.FC = () => {
  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 4 }}>
        <Typography variant="h2" component="h1" gutterBottom>
          Welcome to LoL Metrix
        </Typography>
        <Typography variant="body1" paragraph>
          Your one-stop shop for League of Legends statistics and analysis.
        </Typography>
        <Button
          component={RouterLink}
          to="/about"
          variant="contained"
          color="primary"
        >
          Learn more about LoL Metrix
        </Button>
      </Box>
    </Container>
  );
};

export default Home;