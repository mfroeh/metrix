import React from 'react';
import { Typography, Container, Box, Avatar, Card, CardContent, List, ListItem, ListItemText, Divider } from '@mui/material';
import EmojiEventsIcon from '@mui/icons-material/EmojiEvents';
import { useParams } from 'react-router-dom';

type Summoner = {
    puuid: string;
    profileIconId: number;
    summonerLevel: number;
    name: string;
    tag: string;
};

const Summoner: React.FC = () => {
    const { summonerName, tag } = useParams<{ summonerName: string; tag: string }>();
    const [summonerData, setSummonerData] = React.useState<Summoner | null>(null);

    React.useEffect(() => {
        const fetchSummonerData = async () => {
            try {
                const response = await fetch('/api/v1/summoner', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ name: summonerName, tag: tag }),
                });
                if (!response.ok) {
                    throw new Error('Failed to fetch summoner data');
                }
                const data = await response.json();
                setSummonerData(data["summoner"]);
                console.log("Summoner data:", data);
            } catch (error) {
                console.error('Error fetching summoner data:', error);
            }
        };

        if (summonerName && tag) {
            fetchSummonerData();
        }
    }, [summonerName, tag]);

    if (!summonerData) {
        return <Typography>Loading...</Typography>;
    }

    return (
        <Container maxWidth="sm">
            <Box sx={{ my: 4, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                <Avatar
                    src={`http://ddragon.leagueoflegends.com/cdn/13.10.1/img/profileicon/${summonerData.profileIconId}.png`}
                    sx={{ width: 100, height: 100, mb: 2 }}
                    alt={`${summonerData.name}'s profile icon`}
                />
                <Typography variant="h4" component="h1" gutterBottom>
                    {summonerData.name}#{summonerData.tag}
                </Typography>
                <Typography variant="subtitle1">
                    Level: {summonerData.summonerLevel}
                </Typography>
            </Box>
        </Container>
    );
};

export default Summoner;
