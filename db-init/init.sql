CREATE TABLE Estabelecimento (
	Id serial primary key,
	Nome varchar(50),
	RazaoSocial varchar(50),
	Endereco varchar(100),
	Estado varchar(20),
	Cidade varchar(20),
	Cep varchar(15),
	NumEstabelecimento integer unique
);

CREATE TABLE Loja (
	Id serial primary key,
	Nome varchar(50),
	RazaoSocial varchar(50),
	Endereco varchar(100),
	Estado varchar(20),
	Cidade varchar(20),
	Cep varchar(15),
	NumLoja integer unique,
	IdEstabelecimento integer
);

alter table Loja
add constraint fk_loja
foreign key (IdEstabelecimento)
references Estabelecimento (NumEstabelecimento);