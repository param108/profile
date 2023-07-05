import { hasThread } from "@/app/strings"

describe('parsing threads', ()=>{
    it('finds the thread id and sequence from the tweet', () => {
        expect(true).toBeTruthy()
        const d = hasThread(`#display #thread:a8f4a2c5-d55c-484e-909a-4aff1c382263:9
Hello World.
`)
        expect(d.length).toBe(1);
        expect(d[0].id).toBe("a8f4a2c5-d55c-484e-909a-4aff1c382263");
        expect(d[0].seq).toBe(9);

    })

    it('finds multiple thread id and sequence from tweet', () => {
        expect(true).toBeTruthy()
        const d = hasThread(`#display #thread:a8f4a2c5-d55c-484e-909a-4aff1c382263:9 #thread:a8f4a2c5-d55c-484e-909a-4aff1c382256:10
Hello World.
`)
        expect(d.length).toBe(2);
        expect(d[0].id).toBe("a8f4a2c5-d55c-484e-909a-4aff1c382263");
        expect(d[0].seq).toBe(9);
        expect(d[1].id).toBe("a8f4a2c5-d55c-484e-909a-4aff1c382256");
        expect(d[1].seq).toBe(10);
    })

    it('returns empty array if no threads found', ()=> {
        const d = hasThread(`#display
Hello World.
`)
        expect(d.length).toBe(0);
    })
})
