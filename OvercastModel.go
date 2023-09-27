package main

import "encoding/xml"

type Episode struct {
	Text            string `xml:",chardata"`
	Type            string `xml:"type,attr"`
	OvercastId      string `xml:"overcastId,attr"`
	PubDate         string `xml:"pubDate,attr"`
	Title           string `xml:"title,attr"`
	URL             string `xml:"url,attr"`
	OvercastUrl     string `xml:"overcastUrl,attr"`
	EnclosureUrl    string `xml:"enclosureUrl,attr"`
	UserUpdatedDate string `xml:"userUpdatedDate,attr"`
	UserDeleted     string `xml:"userDeleted,attr"`
	Played          string `xml:"played,attr"`
	Progress        string `xml:"progress,attr"`
}

type Podcast struct {
	Text               string `xml:",chardata"`
	Type               string `xml:"type,attr"`
	Title              string `xml:"title,attr"`
	Smart              string `xml:"smart,attr"`
	Sorting            string `xml:"sorting,attr"`
	ExcludePodcastIds  string `xml:"excludePodcastIds,attr"`
	IncludePodcastIds  string `xml:"includePodcastIds,attr"`
	PriorityPodcastIds string `xml:"priorityPodcastIds,attr"`
	IncludeEpisodeIds  string `xml:"includeEpisodeIds,attr"`
	OvercastId         string `xml:"overcastId,attr"`
	AttrText           string `xml:"text,attr"`
	XmlUrl             string `xml:"xmlUrl,attr"`
	HtmlUrl            string `xml:"htmlUrl,attr"`
	OvercastAddedDate  string `xml:"overcastAddedDate,attr"`
	Subscribed         string `xml:"subscribed,attr"`
	Episodes            []Episode  `xml:"outline"`
}

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Head    struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
	} `xml:"head"`
	Body struct {
		Text    string `xml:",chardata"`
		Outline []struct {
			Text     string `xml:",chardata"`
			AttrText string `xml:"text,attr"`
			Podcasts  []Podcast  `xml:"outline"`
		} `xml:"outline"`
	} `xml:"body"`
} 

