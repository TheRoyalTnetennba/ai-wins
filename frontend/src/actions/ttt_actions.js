type TTTState struct {
    User *datastore.Key
    Marker string
    Started time.Time
    Board [][]string
    Key *datastore.Key `datastore:"__key__"`
}