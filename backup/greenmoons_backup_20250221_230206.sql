PGDMP                      }         
   greenmoons    17.2    17.2                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false                       1262    21066 
   greenmoons    DATABASE     |   CREATE DATABASE greenmoons WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Thai_Thailand.874';
    DROP DATABASE greenmoons;
                     postgres    false            �            1259    21077    breeds    TABLE       CREATE TABLE public.breeds (
    id text NOT NULL,
    name_th text NOT NULL,
    name_en text NOT NULL,
    short_name character varying(30) NOT NULL,
    remark character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.breeds;
       public         heap r       postgres    false                      0    21077    breeds 
   TABLE DATA           b   COPY public.breeds (id, name_th, name_en, short_name, remark, created_at, updated_at) FROM stdin;
    public               postgres    false    217   #       �           2606    21083    breeds breeds_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.breeds
    ADD CONSTRAINT breeds_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.breeds DROP CONSTRAINT breeds_pkey;
       public                 postgres    false    217               <  x��UMk�@=���-%AZ��mem�Œ\V��?�ǤB��{)��%Ł�V�fJgG�q�� X�@ｙ7o=��<|��ѿMqn�w:�Mqm���F/S\P��h<�2�x T(D���,O�' �e�#�1�a�睰�q�޹�?v���O<,��e���p���Z[��'Q[���B�1O�@�I�4x7��s�Q��B�Gbv�mB��<̈́�60�]�>#�)�՗{�D|<�z���� ��ni�����V���"U�iA�Њ�FJ�����_��h���'\!�G�3A��|�ù��u�5��".S<5��>��o�+��C�g�P��IV�6�7�0`(���uID�C-�� 2	gvE�2��(ټ՞oc�tT	}W�V4x�)$fC�$L[�ڈz �	-�-!��,OeQ�$`=^�>U��W�DX�n��d��� x	Y��jt�?^!�L�B 1�������-��t2S>	�is��5�2Iu�����UHe�$��-xed�niv��)d�~jxI���ل˭)�FԵa��6��7y��cy*E�KM`>w:����     