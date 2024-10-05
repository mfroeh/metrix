import React from 'react';
import { Container, Typography, List, ListItem, ListItemText, Paper } from '@mui/material';

const About: React.FC = () => {
  return (
    <Container maxWidth="md">
      <Paper elevation={3} sx={{ padding: 3, marginTop: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          About LoL Metrix
        </Typography>
        <Typography variant="body1" paragraph>
          LoL Metrix is a comprehensive platform for League of Legends statistics and analysis.
        </Typography>
        <Typography variant="body1" paragraph>
          Our goal is to provide players with valuable insights to improve their gameplay and understanding of the game.
        </Typography>
        <List>
          <ListItem>
            <ListItemText primary="In-depth champion statistics" />
          </ListItem>
          <ListItem>
            <ListItemText primary="Match history analysis" />
          </ListItem>
          <ListItem>
            <ListItemText primary="Player performance tracking" />
          </ListItem>
          <ListItem>
            <ListItemText primary="Meta trends and reports" />
          </ListItem>
        </List>
        <Typography variant="body1" paragraph>
          Whether you're a casual player or aspiring pro, LoL Metrix has the tools you need to elevate your League of Legends experience.
        </Typography>
      </Paper>
    </Container>
  );
};

export default About;
