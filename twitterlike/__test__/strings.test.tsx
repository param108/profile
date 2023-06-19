import { tagsToHyperlinks } from '../app/strings'
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
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=param%2Ctweet)`
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
this is an even better [#tweet](https://ui.tribist.com/user/param108/tweets?tags=tweet%2Cparam)`
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
this is an even better [#tweet2](https://ui.tribist.com/user/param108/tweets?tags=tweet%2Cparam%2Ctweet2)`
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
[#param](https://ui.tribist.com/user/param108/tweets?tags=tweet%2Cparam)
[#tent](https://ui.tribist.com/user/param108/tweets?tags=tweet%2Ctent)`
        expect(
            tagsToHyperlinks(
                tweet,
                "https://ui.tribist.com/user/param108/tweets?tags=tweet",
                true
            )).toEqual(expected)
    })

})
