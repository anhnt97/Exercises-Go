Git cơ bản
==========

Tạo một repository mới
----------------------

-   Tạo thư mục chứa project

-   Vào trong thư mục

-   Gõ : git init

Kết nối repo tới github
-----------------------

-   git remote add origin
    [git\@github.com:username/new_repo](git@github.com:username/new_repo)

Clone repo về local
-------------------

-   git clone /đường-dẫn-đến/repository

-   git clone username\@địachỉmáychủ:/đường-dẫn-đến/repository

Tại thư mục repo
----------------

>   Thư mục cục bộ của bạn bao gồm ba "trees" được duy trì bởi git. đầu tiên là
>   Thư Mục Đang Làm Việc (Working Directory) có chứa các tập tin hiện tại. cái
>   thứ hai là Chỉ Mục (Index) đóng vai trò như staging area và cuối cùng là
>   HEAD trỏ đến commit gần đây nhất của bạn.

Add tệp tin vào Index
---------------------

-   git add tệp_tin

-   git add \* (add tất cả tệp tin thay đổi)

    1.  Commit tệp tin thay đổi

-   git commit -m "title" –m “body”

-   Bây giờ thì tập tin đã được commit đến HEAD, nhưng chưa phải trên thư mục
    remote.

    1.  Push lên repo trên github

        Thay đổi hiện đang nằm tại HEAD của bản sao cục bộ đang làm việc. Để gửi
        những thay đổi đó đến repository remote :

        git push origin master

        Thay đổi master bằng bất cứ nhánh nào mà muốn đẩy những thay đổi đến.

Cập nhập repo local
-------------------

Git pull origin nhánh

Branch
------

>   Các nhánh (branches) được dùng để phát triển tính năng tách riêng ra từ
>   những nhánh khác. Nhánh master là nhánh "mặc định" khi tạo một repository.
>   Sử dụng các nhánh khác khi đang trong giai đoạn phát triển và merge trở lại
>   nhánh master một khi đã hoàn tất.

Tạo nhánh
---------

>   git branch \<tên_nhánh\>

Chuyển sang 1 nhánh
-------------------

>   git checkout \<tên_nhánh\>

Vừa tạo vừa chuyển sang nhánh đó
--------------------------------

>   git checkout -b \<tên nhánh\>

Trộn 1 nhánh từ nhánh đang hoạt động 
-------------------------------------

>   git merge \<nhánh\>

>   Sau khi thay đổi, phải đánh dấu chúng là đã được trộn (merged) với lệnh

>   git add \<tên-tập-tin\>

>   trước khi trộn các thay đổi, có thể xem trước chúng bằng cách

>   git diff \<nhánh_nguồn\> \<nhánh_mục_tiêu\>

Git rebase
----------

>   Base lại code dựa theo nhánh chính, khi đó commit của trên nhánh mình sẽ bị
>   thay đổi(thay đổi ở đây chủ yếu là commit hash, nếu trường hợp có conflict
>   thì commit đó cũng sẽ thay đổi luôn cả nội dung)

git rebase \<nhánh\>

1.  git rebase –i

    Thường dùng để merge các commit lại thánh 1 commit hoặc sửa 1 số commit

git rebase -i \<commit\>

1.  Tag

    Dùng git tag khi release 1 chức năng hoặc fix 1 bug gì đấy cho production
    Việc này để cho việc quản lý những thay đổi trên production đơn giản hơn.

    1.  Tạo 1 tag

>   git tag –a v1.3.3 –m “Nội dung của tag này ”

show tag
--------

>   git tag

Reset
-----

>   Nếu bạn muốn hủy tất cả thay đổi và commit cục bộ, lấy về (fetch) lịch sử
>   gần đây nhất từ máy chủ và trỏ nhánh master cục bộ vào nó như sau

>   git fetch origin

>   git reset --hard origin/master

Xóa commit nào đó
-----------------

>   git reset --hard \<commit\>

Sửa message của commit trước
----------------------------

git commit --amend
