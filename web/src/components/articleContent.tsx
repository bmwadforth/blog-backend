import React from 'react';
import {IArticle} from '../store/articles/articlesState';
import useArticle from "../hooks/useArticle";
import {useNavigate, useParams} from "react-router-dom";
import {Stack, Divider, Grid, Paper, Typography, IconButton} from "@mui/material";
import moment from 'moment';
import Loading from "./loading";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {ApplicationRoutes} from "../App";
import {Code} from "@bmwadforth-com/armor-ui";
import {analytics} from "../index";
import {logEvent} from "firebase/analytics";
import {Helmet} from "react-helmet";

export interface IArticleTileProps {
    article: IArticle;
}

export default function ArticleContent() {
    const navigate = useNavigate();
    let {articleSlug} = useParams();
    const [loading, article] = useArticle(articleSlug as string);
    const {data: payload} = article;
    if (loading) return <Loading/>;
    if (payload?.content === undefined) return <h1>Error</h1>;

    logEvent(analytics, 'select_content', {
        content_type: 'article',
        content_id: article.data?.articleId
    });

    return (
        <Paper>
            <Helmet>
                <title>bmwadforth.com | Article - {article.data?.title}</title>
            </Helmet>
            <Grid container style={{padding: '1em', height: '100%'}}>
                <Grid item xs={12}>
                    <Stack direction="row" display='flex' alignItems={'center'} spacing={2}>
                        <Typography
                            variant="subtitle1">
                            <IconButton aria-label="delete" onClick={() => navigate(ApplicationRoutes.ARTICLES)}>
                                <ArrowBackIcon/>
                            </IconButton>
                        </Typography>
                        <Typography
                            variant="subtitle1">{`${payload.title} - ${moment(payload.updated).format('LLL')}`}</Typography>
                    </Stack>
                </Grid>
                <Divider/>
                <Grid item xs={12}>
                    <Code data={payload.content} showLineNumbers/>
                </Grid>
            </Grid>
        </Paper>
    )
}