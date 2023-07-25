import { addThread, tagsToHyperlinks } from '../app/strings'
describe('tweets formatting', ()=>{
    it('renders no tags correctly', ()=>{
        let tweet = `
this is a great tweet with no tags`;
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet",
                true
            )).toEqual(tweet)
    })

    it('renders tags correctly when baseURL doesnt contain tags parameter', ()=> {
        let tweet = `
this is an even better #tweet`
        let expected = `
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=param",
                true
            )).toEqual(expected)
    })

    it('renders tags correctly when baseURL does contain tags parameter', ()=> {
        let tweet = `
this is an even better #tweet`
        let expected = `
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet,param",
                true
            )).toEqual(expected)
    })

    it('renders tags correctly when baseURL contains tags with different tag', ()=> {
        let tweet = `
this is an even better #tweet2`
        let expected = `
this is an even better [#tweet2](https://ui.tribist.com/user/param108/tweets?tags=tweet2)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet,param",
                true
            )).toEqual(expected)
    })

    it('renders tags correctly with multiple', ()=> {
        let tweet = `
this is an even better #tweet
#param
#tent`
        let expected = `
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet)
[#param](https://ui.tribist.com/user/param108/tweets?tags=param)
[#tent](https://ui.tribist.com/user/param108/tweets?tags=tent)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet",
                true
            )).toEqual(expected)
    })

    it('renders tags correctly with repeats', ()=> {
        let tweet = `
this is an even better #tweet
#param
#tweet`
        let expected = `
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet)
[#param](https://ui.tribist.com/user/param108/tweets?tags=param)
[#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet",
                true
            )).toEqual(expected)
    })

    it('adds the thread to tweet without commandline', ()=>{
        let tweet = `Hello Hello Hello
`;
        expect(
            addThread(tweet, "daedf35f-0ff8-4c23-86aa-98213dead4c4")).toEqual(
                `#thread:daedf35f-0ff8-4c23-86aa-98213dead4c4:0
Hello Hello Hello
`
            );
    })

    it('adds the thread to tweet with commandline', ()=>{
        let tweet = `#thread:daedf35f-0ff8-4c23-86aa-98213dead4c4:0
Hello Hello Hello`
        expect(
            addThread(tweet, "0a015aee-4922-4f1c-afef-b321ecd835f7")).toEqual(
                `#thread:daedf35f-0ff8-4c23-86aa-98213dead4c4:0 #thread:0a015aee-4922-4f1c-afef-b321ecd835f7:0
Hello Hello Hello`
            )
    })
})
