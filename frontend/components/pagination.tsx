import React from 'react';

interface IPagination {
    totalPosts: number,
    postsPerPage: number,
    setCurrentPage: Function,
}

const Pagination: React.FC<IPagination> = ({ totalPosts, postsPerPage, setCurrentPage }: { totalPosts: number, postsPerPage: number, setCurrentPage: Function}) => {
    let pages = [];

    for (let i = 1; i <= Math.ceil(totalPosts/postsPerPage); i++) {
        pages.push(i);
    }

    return (
        <div className='pagination'>
            <div className="container">
                <div className="row justify-content-center mt-4">
                    <div className={`col-${Math.ceil(totalPosts/postsPerPage)}`}>
                        <nav aria-label="pagination">
                            <ul className="pagination">
                                {pages.map((page, i) => {
                                    return <li className="page-item" key={i}><a className="page-link" href="#" onClick={() => setCurrentPage(page)}>{page}</a></li>;
                                })}
                            </ul>
                        </nav>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Pagination;