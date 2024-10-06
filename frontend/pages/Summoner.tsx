import React from 'react';
import { Typography, Container, Box, Avatar, Card, CardContent, List, ListItem, ListItemText, Divider } from '@mui/material';
import EmojiEventsIcon from '@mui/icons-material/EmojiEvents';
import { useParams } from 'react-router-dom';

type League = {
    queueType: string;
    tier: string;
    rank: number;
    wins: number;
    losses: number;
    leaguePoints: number;
    ratedRating: number;
    hotStreak: boolean;
    veteran: boolean;
    freshBlood: boolean;
};

type Summoner = {
    puuid: string;
    profileIconId: number;
    summonerLevel: number;
    name: string;
    tag: string;
    leagues: League[];
};

const LeagueInfo: React.FC<{ league: League }> = ({ league }) => {
    const getTierIcon = (tier: string) => {
        const lowerTier = tier.toLowerCase();
        return `https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/${lowerTier}.png`;
    };

    const getRomanNumeral = (rank: number) => {
        const romanNumerals = ['I', 'II', 'III', 'IV'];
        return romanNumerals[rank - 1] || '';
    };

    return (
        <Box sx={{ mt: 2, p: 2, border: 1, borderColor: 'grey.300', borderRadius: 2 }}>
            <Typography variant="h6">{league.queueType}</Typography>
                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <img 
                        src={getTierIcon(league.tier)} 
                        alt={`${league.tier} icon`} 
                        style={{ width: '24px', height: '24px', marginRight: '8px' }}
                    />
                    <Typography>
                        {league.tier} {getRomanNumeral(league.rank)}
                    </Typography>
                </Box>
                <Typography>Wins: {league.wins} / Losses: {league.losses}</Typography>
                {league.leaguePoints !== undefined && (
                    <Typography>League Points: {league.leaguePoints}</Typography>
                )}
                {league.ratedRating !== undefined && (
                    <Typography>Rated Rating: {league.ratedRating}</Typography>
                )}
                <Typography>
                    {league.hotStreak && '🔥 Hot Streak '}
                    {league.veteran && '🏅 Veteran '}
                    {league.freshBlood && '🆕 Fresh Blood '}
                </Typography>
        </Box>
    );
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
                {summonerData.leagues.map((league) => <LeagueInfo key={league.queueType} league={league} />)}
            </Box>
        </Container>
    );
};

export default Summoner;
