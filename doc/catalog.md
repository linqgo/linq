# Functional Catalog

The following function summary groups all functions into major categories.
Aside from the first group, **Query methods**, all functions exist as free
functions. Due to limitations in Go generics, only some functions are also
available as methods.

### Syntax guide

- `[...]` - optional
- `{...}` - non-optional (used for syntactic grouping)
- `Ⓜ️` - implemented both as a free function and as a method of `Query[T]`
- <code><strong>⁺</strong></code> - new capability, not ported from .Net
- `term1,term2...` - one of term1, term2…

## github.com/linqgo/linq

<table><tbody><tr>
<td>
    <h4>Query methods</h4>
    <ul>
        <li><code>Enumerator</code></li>
        <li><code>OneShot</code></li>
    </ul>
    <h4>Construct</h4>
    <ul>
        <li><code>From[ByteReader,Channel,Map,RuneReader,[Scanner][String]]</code></li>
        <li><code>Iota[1,2,3]</code> (equivalent to <code>Enumerable.Range</code> in .Net)</li>
        <li><code>NewQuery</code></li>
        <li><code>None</code></li>
        <li><code>Pipe</code></li>
        <li><code>Repeat[Forever<strong>⁺</strong>]</code></li>
    </ul>
    <h4>Convert to Go types</h4>
    <ul>
        <li><code>[Must]ToMap[KV<strong>⁺</strong>]</code></li>
        <li><code>Ⓜ️ToSlice</code></li>
        <li><code>ToString<strong>⁺</strong></code></li>
    </ul>
    <h4>Aggregation</h4>
    <ul>
        <li><code>Ⓜ️Aggregate[Seed<strong>¹</strong>]</code></li>
        <li><code>Average,Sum</code></li>
        <li><code>Ⓜ️Count[Limit[True]<strong>⁺</strong>],Ⓜ️FastCount<strong>⁺</strong></code></li>
        <li><code>Strings{Join,CommaAnd}<strong>⁺</strong></code></li>
    </ul>
    <h4>Element selection</h4>
    <ul>
        <li><code>Ⓜ️[Fast]ElementAt</code></li>
        <li><code>Ⓜ️First[Comp<strong>⁺</strong>]</code></li>
        <li><code>Ⓜ️Last</code></li>
        <li><code>{Max,Min}[By]</code></li>
        <li><code>Ⓜ️Single<strong>⁺</strong></code></li>
    </ul>
    <h4>Predicate</h4>
    <ul>
        <li><code>Ⓜ️All</code></li>
        <li><code>Ⓜ️Any</code></li>
        <li><code>Contains</code></li>
        <li><code>Ⓜ️Empty</code></li>
        <li><code>SequenceEqual[Ⓜ️Eq<strong>⁺</strong>],Sequence{Greater,Less}[Ⓜ️Comp]<strong>⁺</strong></code></li>
        <li><code>Ⓜ️[Fast]{Longer,Shorter}<strong>⁺</strong></code></li>
    </ul>
    <h4>Compose</h4>
    <ul>
        <li><code>Ⓜ️Append,Ⓜ️Prepend</code></li>
        <li><code>Ⓜ️Concat</code></li>
    </ul>
    <h4>Transform</h4>
    <ul>
        <li><code>Index[From]<strong>⁺</strong></code></li>
        <li><code>Pairwise<strong>⁺</strong></code></li>
        <li><code>Ⓜ️<strong>¹</strong>Select</code></li>
        <li><code>Select[Keys<strong>⁺</strong>,Many,Values<strong>⁺</strong>]</code></li>
        <li><code>Unzip[KV]<strong>⁺</strong></code></li>
        <li><code>Zip[KV<strong>⁺</strong>]</code></li>
    </ul>
</td>
<td>
    <h4>Filter</h4>
    <ul>
        <li><code>Distinct[By]</code></li>
        <li><code>Ⓜ️Every[From]<strong>⁺</strong></code></li>
        <li><code>OfType</code></li>
        <li><code>Ⓜ️Sample[Seed]<strong>⁺</strong></code></li>
        <li><code>Ⓜ️Skip[Last,While]</code></li>
        <li><code>Ⓜ️Take[Last,While]</code></li>
        <li><code>Ⓜ️Where</code></li>
    </ul>
    <h4>Rearrange</h4>
    <ul>
        <li><code>Ⓜ️Reverse</code></li>
        <li><code>{Order,Then}[By,Ⓜ️Comp][Desc]</code></li>
    </ul>
    <h4>Group and ungroup</h4>
    <ul>
        <li><code>Chunk[Slices]</code></li>
        <li><code>Flatten[Slices]<strong>⁺</strong></code></li>
        <li><code>GroupBy[Select][Slices]</code></li>
        <li><code>GroupJoin</code></li>
        <li><code>SelectMany</code></li>
        <li><code>Slide[All,Fixed,Time]<strong>⁺</strong>,Delta<strong>⁺</strong></code></li>
    </ul>
    <h4>Set and relational operations</h4>
    <ul>
        <li><code>Except[By]</code></li>
        <li><code>Join</code></li>
        <li><code>Intersect[By]</code></li>
        <li><code>PowerSet<strong>⁺</strong></code></li>
        <li><code>Union</code></li>
    </ul>
    <h4>Miscellaneous helpers</h4>
    <ul>
        <li><code>Array<strong>⁺</strong>,Getter<strong>⁺</strong></code></li>
        <li><code>Drain<strong>⁺</strong></code></li>
        <li><code>[Not]Equal<strong>⁺</strong>,Less<strong>⁺</strong>,Greater<strong>⁺</strong></code></li>
        <li><code>False<strong>⁺</strong>,True<strong>⁺</strong>,Zero<strong>⁺</strong></code></li>
        <li><code>Identity<strong>⁺</strong></code></li>
        <li><code>Key<strong>⁺</strong>,Value<strong>⁺</strong></code></li>
        <li><code>{Longer,Shorter}{Slice,Map}<strong>⁺</strong></code></li>
        <li><code>Maybe[Ⓜ️<strong>¹</strong>FlatMap]<strong>⁺</strong>,Some<strong>⁺</strong>,No<strong>⁺</strong></code></li>
        <li><code>Not<strong>⁺</strong></code></li>
        <li><code>Pointer<strong>⁺</strong>,Deref<strong>⁺</strong></code></li>
        <li><code>Scanner<strong>⁺</strong></code></li>
        <li><code>SwapArgs<strong>⁺</strong></code></li>
    </ul>
    <h4>Miscellaneous</h4>
    <ul>
        <li><code>Ⓜ️DefaultIfEmpty</code></li>
        <li><code>Ⓜ️Memoize<strong>⁺</strong></code></li>
        <li><code>NewKV</code></li>
    </ul>
</td>
</tr></tbody></table>

Notes:<br/>
<strong>¹</strong> The method is more restrictive than the function due to limitations with Go generics.<br/>

## github.com/linqgo/linq/stats

<table><tbody><tr>
<td>
    <h4>Math aggregation</h4>
    <ul>
        <li><code>[Acc<strong>⁺</strong>]Mean</code></li>
        <li><code>[Acc]GeometricMean<strong>⁺</strong></code></li>
        <li><code>[Acc]HarmonicMean<strong>⁺</strong></code></li>
        <li><code>[Acc]Product<strong>⁺</strong></code></li>
        <li><code>[Acc<strong>⁺</strong>]Sum</code></li>
    </ul>
</td>
</tr></tbody></table>
